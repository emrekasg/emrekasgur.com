const { execSync } = require('child_process');
const { processPost, insertOrUpdatePost } = require('./post');

module.exports.handler = async (event) => {
  const repo = event.repo;

  const repoName = Math.random().toString(14).substring(7);

  execSync(`git clone ${repo} /tmp/${repoName}`);
  process.chdir(`/tmp/${repoName}`);

  // keep changed files that exists in ./posts folder in a variable
  const changedFilesCmd = `git diff --name-only HEAD HEAD~1 | grep posts/`;
  const changedFiles = execSync(changedFilesCmd).toString().split("\n");

  const posts = [];

  // check if there's empty string in the array and remove it
  const emptyStringIndex = changedFiles.indexOf("");
  if (emptyStringIndex > -1) {
    changedFiles.splice(emptyStringIndex, 1);
  }

  // if there is no change in ./posts folder, exit
  if (changedFiles.length === 0) {
    return {
      statusCode: 200,
      body: JSON.stringify(
        {
          message: 'No changes in ./posts folder',
          input: event,
        },
        null,
        2
      ),
    };
  }

  // if there is a change in ./posts folder, read the changed files one by one
  changedFiles.forEach((file) => {
    const post = processPost(file);
    posts.push(post);
  });


  for (let i = 0; i < posts.length; i++) {
    const post = posts[i];

    await insertOrUpdatePost(post);
  }

  return {
    statusCode: 200,
    body: JSON.stringify(
      {
        message: 'Go Serverless v3.0! Your function executed successfully!s',
        input: event,
      },
      null,
      2
    ),
  };
};