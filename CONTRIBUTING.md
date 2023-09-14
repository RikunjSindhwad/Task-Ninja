# Contribution Guidelines

We welcome and appreciate contributions from the community to help improve TaskNinja. Whether it's fixing a bug, adding a new feature, or enhancing documentation, your contributions make TaskNinja better for everyone. Before you start contributing, please read and follow these guidelines:

## Code of Conduct

Please review and adhere to our [Code of Conduct](https://github.com/RikunjSindhwad/Task-Ninja/wiki/Code-of-Conduct). We expect respectful and inclusive behavior from all contributors.

## Getting Started

Before you begin contributing, make sure you have TaskNinja installed and understand how it works. Refer to our [Getting Started](https://github.com/RikunjSindhwad/Task-Ninja/wiki/Getting-Started) guide for installation instructions.

## How to Contribute

1. **Fork the Repository:** Click the "Fork" button on the top right of the repository's page to create a copy in your GitHub account.

2. **Clone Your Fork:** Clone the forked repository to your local machine using `git clone`.

```bash
   git clone https://github.com/RikunjSindhwad/Task-Ninja.git
```
3. **Create a Branch:** Create a new branch for your contribution.

```bash
  git checkout -b feature/your-feature
```
4. **Make Changes:** Make your desired changes in the codebase. Ensure your code adheres to the existing coding style and standards.

5. **Commit Changes:** Commit your changes with a clear and concise commit message.

```bash
git commit -m "Add your commit message here"
```
6. **Push to Your Fork:** Push your changes to your forked repository on GitHub.

```bash
   git push origin feature/your-feature
```
7. **Create a Pull Request:** Go to the TaskNinja repository on GitHub, and you should see a "Compare & pull request" button. Click it to open a new pull request.
   - Make sure your base branch is `main` if you're contributing to the main repository.
   - Ensure that your branch with changes is selected as the compare branch.

8. **Describe Your Pull Request:** Provide a clear and concise description of your changes in the pull request. Explain what the changes do and why they are necessary.

   - Use descriptive and meaningful titles and descriptions.
   - If your pull request addresses or fixes an existing issue, reference the issue number in the description using GitHub's syntax (e.g., "Fixes #123").

9. **Review and Address Feedback:** Collaborate with reviewers to address any feedback or changes requested during the code review process.

   - Be open to constructive criticism and improvements suggested by reviewers.
   - Make the necessary changes and update your pull request accordingly.

10. **Wait for Approval:** Once your pull request is approved, it will be merged into the main repository.

    - A project maintainer will review your changes and merge them if they meet the project's standards and requirements.

11. **Stay Updated:** Keep your fork up to date with the upstream repository to ensure your code remains compatible with the latest changes.

 ```bash
    git remote add upstream https://github.com/robenstark/TaskNinja.git
    git fetch upstream
    git checkout main
    git merge upstream/main
    git push origin main
```

12. **Celebrate:** Congratulations! You've successfully contributed to TaskNinja. Thank you for your valuable contribution to the project.
