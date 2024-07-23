# vcs

A simple version control system written in Go with the Cobra library for creating CLI applications in Go. Inspired by [Build Your Own VCS by Ryan Heathcote](https://ryanheathcote.com/git/build-your-own-vcs).

# TODO

- [ ] fix status showing <file> to be committed after it is already committed

- [ ] implement unstaging changes
- [ ] implement diff

- [ ] implement better pattern handling in `.vcsignore`
- [ ] implement better pattern handling in `add`
    - e.g. `go run main.go add .` adds a file named `"."` rather than all unstaged files

- [ ] handle multiple files being added at the same time
    ```
    go run main.go add LICENSE README.md
    File added: LICENSE
    ```

- [ ] implement commit history (git log)
