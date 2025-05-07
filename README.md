# GitHub User Activity 

Use GitHub API to fetch user activity and display it in the terminal. 

In this project works with a simple command line interface (CLI) to fetch the recent activity of a GitHub user and display it in the terminal. The purpose of this project is practice  programming skills, including working with APIs, handling JSON data, and building a simple CLI application.

Requirements
The application should run from the command line, accept the GitHub username as an argument, fetch the user’s recent activity using the GitHub API, and display it in the terminal. The user should be able to:

- Provide the GitHub username as an argument when running the CLI. 
- Fetch the recent activity of the specified GitHub user using the GitHub API using the following endpoint to fetch the user’s activity: 
    - https://api.github.com/users/<username>/events
- Display the fetched activity in the terminal. 
- Handle errors gracefully, such as invalid usernames or API failures.
- Use a programming language of your choice to build this project.
- Do not use any external libraries or frameworks to fetch the GitHub activity.

Future improvements:
Filtering the activity by event type, displaying the activity in a more structured format, or caching the fetched data to improve performance. You can also explore other endpoints of the GitHub API to fetch additional information about the user or their repositories.
