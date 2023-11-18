const fs = require("fs");
const path = require("path");
const mysql = require('mysql2/promise');

const db = process.env.DATABASE_URL || "";

module.exports.insertOrUpdatePost = async function (post) {
    const connection = await mysql.createConnection(db);
    try {
        await connection.beginTransaction();
        const checkPostLinkQuery = 'SELECT id FROM posts WHERE post_link = ?';
        const [postRow] = await connection.query(checkPostLinkQuery, [post.post_link]);

        let postId;

        if (postRow.length === 0) {
            const insertPostQuery = 'INSERT INTO posts (post_link, tag, visible) VALUES (?, ?)';
            const [rows] = await connection.query(insertPostQuery, [post.post_link, post.tag, post.visible]);
            postId = rows.insertId;
        } else {
            postId = postRow[0].id;
            const updatePostQuery = 'UPDATE posts SET tag = ?, visible = ? WHERE id = ?';
            await connection.query(updatePostQuery, [post.tag, post.visible, postId]);
        }

        const checkLanguageQuery = 'SELECT id FROM post_contents WHERE post_id = ? AND lang = ?';
        const [languageRow] = await connection.query(checkLanguageQuery, [postId, post.language]);

        if (languageRow.length === 0) {
            const insertPostContentQuery = 'INSERT INTO post_contents (post_id, lang, content, brief, title) VALUES (?, ?, ?, ?, ?)';
            await connection.query(insertPostContentQuery, [postId, post.language, post.content, post.brief, post.title]);
        } else {
            const updatePostContentQuery = 'UPDATE post_contents SET content = ?, brief = ?, title = ? WHERE id = ?';
            await connection.query(updatePostContentQuery, [post.content, post.brief, post.title, languageRow[0].id]);
        }

        await connection.commit();
    } catch (error) {
        await connection.rollback();
        throw error;
    }
}

/*
    content always starts like:
        [PostLink] = "{{ post_link }}"
        [PostTitle] = "{{ title }}"
        [Brief] = "{{ brief }}"
        [Language] = "{{ language_code }}"
        [Tag] = "{{ tag }}"
        [Visible] = true/false
        
    We need to parse those values and keep them in variables
*/
module.exports.processPost = function (file) {
    const content = fs.readFileSync(path.join(process.cwd(), file), "utf8");
    const lines = content.split("\n");

    const postMap = {
        "[PostLink]": "post_link",
        "[PostTitle]": "title",
        "[Brief]": "brief",
        "[Language]": "language",
        "[Tag]": "tag",
        "[Visible]": "visible",
    };

    const postMapLength = Object.keys(postMap).length;
    for (let i = 0; i < postMapLength; i++) {
        const line = lines[i];
        const [key, value] = line.split(" = ");

        // if the line is not in the postMap, skip it
        if (!postMap[key]) {
            continue;
        }

        if (key === "[Visible]") {
            postMap[postMap[key]] = value === "true" ? 1 : 0;
            continue;
        }

        postMap[postMap[key]] = value.replace(/"/g, "");
    }

    // keep the rest of the content as the body
    const body = lines.slice(postMapLength+1).join("\n");
    postMap["content"] = body;

    return postMap;
}