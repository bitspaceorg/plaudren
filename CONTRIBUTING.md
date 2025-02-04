# Contributing

**Plaudren** uses GitHub to manage reviews of pull requests.
* If you have a trivial fix or improvement, go ahead and create a pull request,
  addressing (with `@...`) a suitable maintainer of this repository (see
  [MAINTAINERS.md](MAINTAINERS.md)) in the description of the pull request.

* If you plan to do something more involved, first discuss your ideas
  on our [Discord server](https://discord.com/invite/Z3PdY9YmHZ).
  This will avoid unnecessary work and surely give you and us a good deal
  of inspiration. 

* Make sure to refer Formatting and style_ section of Peter Bourgon's [Go: Best
  Practices for Production
  Environments](https://peter.bourgon.org/go-in-production/#formatting-and-style).


## Steps to Contribute

Should you wish to work on an issue, please claim it first by commenting on the GitHub issue that you want to work on it. This is to prevent duplicated efforts from contributors on the same issue.

Please check the `good-first-issue` label to find issues that are good for getting started. If you have questions about one of the issues, with or without the tag, please comment on them and one of the maintainers will clarify it. For a quicker response, contact us over [Discord](https://discord.com/invite/Z3PdY9YmHZ).

For quickly compiling and testing your changes do:

```bash
go test -v ./... # make sure all the tests pass.
```

All our issues are regularly tagged so that you can also filter down the issues involving the components you want to work on. 

## Pull Request Checklist

* Branch from the `dev` branch and, if needed, rebase to the current `dev` branch before submitting your pull request. If it doesn't merge cleanly with `dev` you may be asked to rebase your changes.

* Commits should be as small as possible, while ensuring that each commit is correct independently (i.e., each commit should compile and pass tests).

* If your patch is not getting reviewed or you need a specific person to review it, you can @-reply a reviewer asking for a review in the pull request or a comment, or you can ask for a review on the [Discord Server](https://discord.com/invite/Z3PdY9YmHZ)
 
* Add tests relevant to the fixed bug or new feature.

## Dependency management
- **plaudren** is a zero dependency project, so refrain from adding any additional dependecies and try to utilize the go standard library.