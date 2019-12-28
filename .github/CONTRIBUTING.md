# Contributing to [twitch-emotes-downloader](https://github.com/Guergeiro/twitch-emotes-downloader/)

## Bug Reports
A bug is a *demonstrable problem* that is caused by the code in the repository. Good bug reports are extremely helpful, so thanks!
If you want to report a bug, click [here](https://github.com/Guergeiro/twitch-emotes-downloader/issues/new?assignees=&labels=&template=bug_report.md&title=).

## Feature Requests
Feature requests are welcome. But take a moment to find out whether your idea fits with the scope and aims of the project. It's up to *you* to make a strong case to convince the project developer of the merits of this feature. Please provide as much detail and context as possible. If you want to request a feature, click [here](https://github.com/Guergeiro/twitch-emotes-downloader/issues/new?assignees=&labels=&template=feature_request.md&title=).

## Pull Requests
1. [Fork](https://help.github.com/articles/fork-a-repo/) the project, clone your fork, and configure the remotes:
```bash
# Clone your fork of the repo into the current directory
git clone https://github.com/<your-username>/twitch-emotes-downloader.git
# Navigate to the newly cloned directory
cd twitch-emotes-downloader
# Assign the original repo to a remote called "upstream"
git remote add upstream https://github.com/Guergeiro/twitch-emotes-downloader.git
```
2. If you cloned a while ago, get the latest changes from upstream:
```bash
git checkout develop
git pull upstream develop
```
3. Create a new topic branch (off the main project development branch) to contain your feature, change, or fix:
```bash
git checkout -b <topic-branch-name>
```
4. Locally merge (or rebase) the upstream development branch into your topic branch:
```bash
git pull [--rebase] upstream develop
```
5. Push your topic branch up to your fork:
```bash
git push origin <topic-branch-name>
```
6. [Open a Pull Request](https://help.github.com/articles/about-pull-requests/) with a clear title and description against the `develop` branch.

## Code Guidelines
- Code should follow all [Google Style](https://google.github.io/styleguide/pyguide.html) guidelines.
- Code should be formatted according to [YAPF](https://github.com/google/yapf/).

## License
By submitting a patch, you agree to allow the project owners to license your work under the terms of the [MIT License](https://github.com/Guergeiro/twitch-emotes-downloader/blob/master/LICENSE).
