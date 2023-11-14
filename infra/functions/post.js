const fs = require("fs");
const path = require("path");
const mysql = require('mysql2')

const db = process.env.DATABASE_URL || "";
const connection = mysql.createConnection(db);

module.exports.insertOrUpdatePost = async function (post) {
    connection.connect();
    const checkPostLinkQuery = 'SELECT id FROM posts WHERE post_link = ?';
    const [postRow] = await connection.promise().query(checkPostLinkQuery, [post.post_link]);

    let postId;

    if (postRow.length === 0) {
        const insertPostQuery = 'INSERT INTO posts (post_link) VALUES (?)';
        const [rows] = await connection.promise().query(insertPostQuery, [post.post_link]);
        postId = rows.insertId;
    } else {
        postId = postRow[0].id;
    }

    const checkLanguageQuery = 'SELECT id FROM post_contents WHERE post_id = ? AND lang = ?';
    const [languageRow] = await connection.promise().query(checkLanguageQuery, [postId, post.language]);

    if (languageRow.length === 0) {
        const insertPostContentQuery = 'INSERT INTO post_contents (post_id, lang, content, brief, title) VALUES (?, ?, ?, ?, ?)';
        await connection.promise().query(insertPostContentQuery, [postId, post.language, post.content, post.brief, post.title]);
    } else {
        const updatePostContentQuery = 'UPDATE post_contents SET content = ?, brief = ?, title = ? WHERE id = ?';
        await connection.promise().query(updatePostContentQuery, [post.content, post.brief, post.title, languageRow[0].id]);
    }
}


/*
  content always starts like:
    [PostLink] = "setupping-envoy-sidecar-on-fargate-pods-in-realtime"
    [PostTitle] = "Setupping envoy sidecar, and gathering metrics on Fargates/pods in real-time without App Mesh"
    [Brief] = "In this article we will setup envoy sidecar on Fargate pods, and gather metrics in real-time without App Mesh"
    [Language] = "en"
    
  I need to parse these values and keep them in variables
*/
module.exports.processPost = function (file) {
    const content = fs.readFileSync(path.join(process.cwd(), file), "utf8");

    // split the content by new line
    const lines = content.split("\n");

    let postLink, postTitle, brief, language, body;
    
    // iterate first 5 lines
    lines.forEach((line, index) => {
        const [name, value] = line.split(" = ");
        switch (name) {
            case "[PostLink]":
                postLink = value;
                break;
            case "[PostTitle]":
                postTitle = value;
                break;
            case "[Brief]":
                brief = value;
                break;
            case "[Language]":
                language = value;
                break;
            default:
                break;
        }
    });

    // keep the rest of the content as the body
    body = lines.slice(5).join("\n");

    return {
        post_link: postLink,
        title: postTitle,
        brief,
        language,
        content: body
    };
}