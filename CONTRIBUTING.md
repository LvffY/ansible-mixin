[[_TOC_]]

## Project Structure

In the `cmd/skeletor` directory, you will find a cli built using [spf13/cobra](https://github.com/spf13/cobra). The CLI contains a go file for each basic capability a Mixin should implement:

* build
* schema
* version
* install
* upgrade
* invoke
* uninstall

Each of these command implementations have a corresponding Mixin implementation in the `pkg/skeletor` directory. Each of the commands above is wired into an empty implementation in `pkg/skeletor` that needs to be completed. In order to build a new Mixin, you need to complete these implementations with the relevant technology. For example, to build a [Cloud Formation](https://aws.amazon.com/cloudformation/) mixin, you might implement the methods in `pkg/skeletor` using the [AWS Go SDK](https://docs.aws.amazon.com/sdk-for-go/api/service/cloudformation/).

## Provided capabilities

This skeleton mixin project brings some free capabilities:

### File System Access and Context

Porter provides a [Context](https://porter.sh/src/pkg/context) package that has helpful mechanisms for accessing the File System using [spf13/afero](https://github.com/spf13/afero). This makes it easy to provide mock File System implementations during testing. The Context package also provides a mechanism to encapsualte stdin, stdout and stderr so that they can easily be passed from `cmd/skeletor` code to implementing `pkg/skeletor` code.

### Template and Static Asset Handling

The project already includes [Packr V2](https://github.com/gobuffalo/packr/tree/master/v2) for dealing with static files, such as templates or other content that is best modeled outside of a Go file. You can see an example of this in `pkg/skeletor/schema.go`.

### Basic Schema

The project provides an implementation of the `skeletor schema` command that is mostly functional. To fully implement this for your mixin, you simply need to provide a valid JSON schema. For reference, consult `pkg/skeletor/schema/schema.json`.

### Basic Tests

The project provides some very basic test skeletons that you can use as a starting point for building tests for your mixin.

### Magefile

The project also includes a [Magefile] that is used to build, test, and publish the mixin.

### Publish

You must set the `GITHUB_TOKEN` environment variable with your personal access token in order to use the default publish target.

Publish uploads cross-compiled binaries of your mixin to a GitHub release.
You must set the `PORTER_RELEASE_REPOSITORY` environment variable to your GitHub repository name, e.g. github.com/YOURNAME/YOURREPO.
There is a placeholder in the Publish magefile target where you can set that value.

Create a tag, for example `git tag v0.1.0`, and push it to your repository.
Run `mage XBuildAll Publish` to build your mixin and upload the binaries to the github release for that tag.
If the commit is not tagged, the release is named "canary".

If you want to generate a mixin feed file (atom.xml), edit the Publish magefile target, uncomment out the rest of the function, and set the `PORTER_PACKAGES_REMOTE` environment variable to a repository where the atom.xml file should be committed.
For example, Porter uses github.com/getporter/packages for publishing our mixin feed.

[Magefile]: https://magefile.org
