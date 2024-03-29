This document describes the way you can contribute to the project.

Ways to Help Out
================

Thank you for your contribution! Please **open an issue first** (or add a
comment to an existing issue) if you plan to work on any code or add a new
feature. This way, duplicate work is prevented and we can discuss your ideas
and design first.

There are several ways you can help us out. First of all code contributions and
bug fixes are most welcome. However, even "minor" details such as fixing spelling
errors, improving the documentation or pointing out usability issues are a great
help also.

The project uses the GitHub infrastructure (see the
[project page](https://github.com/melnyk/go-concept)) for all related discussions.

If you want to find an area that currently needs improving have a look at the
open issues listed on the
[issues](https://github.com/melnyk/go-concept/issues) page](https://github.com/melnyk/go-concept/issues). This is also the place
for discussing enhancement to the tools.


Reporting Bugs
==============

You've found a bug? Thanks for letting us know so we can fix it! It is a good
idea to describe in detail how to reproduce the bug (when you know how), what
environment was used and so on. Please tell us at least the following things:

 * What's the version of the module you used?
 * What commands did you execute to get to where the bug occurred?
 * What did you expect?
 * What happened instead?
 * Are you aware of a way to reproduce the bug?

Remember, the easier it is for us to reproduce the bug, the earlier it will be
corrected!

Development Environment
=======================

The repository contains the code written for the project in the root directory.

Make sure you have the minimum required Go version installed. Clone the repo
(without having `$GOPATH` set) and `cd` into the directory:

    $ unset GOPATH
    $ git clone https://github.com/melnyk/go-concept
    $ cd go-uuid

Then use the `go build` tool to build the module


Providing Patches
=================

You have fixed an annoying bug or have added a new feature? Very cool! Let's
get it into the project! The workflow we're using is also described on the
[GitHub Flow](https://guides.github.com/introduction/flow/) website, it boils
down to the following steps:

 0. If you want to work on something, please add a comment to the issue on
    GitHub. For a new feature, please add an issue before starting to work on
    it, so that duplicate work is prevented.

 1. Next, fork our project on GitHub if you haven't done so already.

 2. Clone your fork of the repository locally and **create a new branch** for
    your changes. If you are working on the code itself, please set up the
    development environment as described in the previous section.

 3. Commit your changes to the new branch as fine-grained as possible, as
    smaller patches, for individual changes, are easier to discuss and merge.

 4. Push the new branch with your changes to your fork of the repository.

 5. Create a pull request by visiting the GitHub website, it will guide you
    through the process. Please [allow edits from maintainers](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/allowing-changes-to-a-pull-request-branch-created-from-a-fork).

 6. You will receive comments on your code and the feature or bug that they
    address. Maybe you need to rework some minor things, in this case, push new
    commits to the branch you created for the pull request (or amend the
    existing commit, use common sense to decide which is better), and they will be
    automatically added to the pull request.

 7. If your pull request changes anything that users should be aware of
    (a bugfix, a new feature, ...) please provide information about these changes.
    It will be used in the announcement of the next stable release. While
    writing, ask yourself: If I were the user, what would I need to be aware
    of with this change?

 8. Once your code looks good and passes all the tests, we'll merge it. Thanks
    a lot for your contribution!

Please provide the patches for each bug or feature in a separate branch and
open up a pull request for each, as this simplifies discussion and merging.

The project uses the `gofmt` tool for Go source indentation, so please
run

    gofmt -w **/*.go

in the project root directory before committing. For each Pull Request, the
formatting is tested with `gofmt` for the latest stable version of Go.

For each pull request, several different systems run the integration tests on
Linux, macOS and Windows. We won't merge any code that does not pass all tests
for all systems, so when a test fails, try to find out what's wrong and fix
it. If you need help on this, please leave a comment in the pull request, and
we'll be glad to assist. Having a PR with failing integration tests is nothing
to be ashamed of. In contrast, that happens regularly for all of us. That's
what the tests are there for.

Git Commits
-----------

It would be good if you could follow the same general style regarding Git
commits as the rest of the project, this makes reviewing code, browsing the
history and triaging bugs much easier.

Git commit messages to have a very terse summary in the first line of the commit
message, followed by an empty line, followed by a more verbose description or a
List of changed things. For examples, please refer to the excellent [How to
Write a Git Commit Message](https://chris.beams.io/posts/git-commit/).

If you change/add multiple different things that aren't related at all, try to
make several smaller commits. This is much easier to review. Using `git add -p`
allows staging and committing only some changes.

Code Review
===========

The project encourages actively reviewing the code, as it will store
your precious data, so it's common practice to receive comments on provided
patches.

If you are reviewing other contributor's code please consider the following
when reviewing:

* Be nice. Please make the review comment as constructive as possible so all
  participants will learn something from your review.

As a contributor, you might be asked to rewrite portions of your code to make it
fit better into the upstream sources.
