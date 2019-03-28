|              |                                  |
| :----------- | :------------------------------- |
| Feature Name | Suse Organization `Webhooks`       |
| Start Date   | 27.03.2019                       |
| RFC PR       | (leave this empty)               |

# Summary
[summary]: #summary

As a technology company we must use all available technologies to ease our way to produce and deliver software, in this case I am proposing to enable the github `webhooks` at organization level so different teams can generate workflows according to their necessities.

# Motivation
[motivation]: #motivation

- Automate Processes
- Automate Project Management

To help us keep control of our BugSquad in CaaSP team we want to automate a project management tool using github projects, you can see a beta version [here](https://github.com/SUSE/caasp-playground/projects/3), this helps motivate the developer, qa and release engineers to work more and spend less time managing boards.

Since all our work is done in github, we should use the same tool to keep track of the progress; so assigning or un-assigning, labeling or un-labeling an issue, PR should generate the correct actions to track our progress.

To help us do this we need to use an organization project so we can reference multiple repositories in the same project, and to get the correct actions from GitHub we need to enable GitHub `Webhooks` at an organization level.

# Detailed design
[design]: #detailed-design

The idea behind this is that not only CaaSP team will benefit from this but all SUSE organization, therefore I am proposing the following design for the application:

```
├── pkg
│   ├── caasp
│   │   └── caasp.go
│   ├── config
│   │   └── config.go
│   ├── github
│   │   ├── request
│   │   │   ├── request.go
│   │   │   ├── request_test.go
│   │   └── model
│   │       ├── comment.go
│   │       ├── issue.go
│   │       ├── label.go
│   │       ├── note.go
│   │       └── pullrequest.go
│   └── security
│       ├── signature.go
│       └── signature_test.go
├── server
│   └── server.go
├── version
│   └── version.go
├── main.go
```

## Package by Package

### server

The server package contains the HTTP server that will be used to listen for the `Webhooks` from GitHub. The idea is to filter the `Webhooks` based on the repoository they come from so they get prodessed by the correct handler.

### version

This package will display the current running version of the application.

### pkg/github

This package contains the models used to unmarshal the `Webhooks`.
Here is also the handler to send request to GitHub.

### pkg/config

This package contains the model for the config file, which is parsed at the start of the application. (Description [here](##config-file-description))

### pkg/security

This package to verifies that all the request sent to the application match the secret key that was set in GitHub. If this security verification fails, no further processing is done. This prevent falling for fake requests sent to the application.

### pkg/caasp

This is the package that will take care of all the actions that the `CaaSP team` need to perform when receiving a `Webhook` we are interested on processing.

### pkg/\<team\>

The idea is that any new team that whats to process `Webhooks` can add the actions in a new package. This package will normally contain a Handler that will be used in the server to resolve actions depending of the repository the `Webhook` comes from.

## Config file description

```
---
server:
  address: ":4321"
  readTimeout: 10
  writeTimeout: 10
github:
  apiURL: "https://api.github.com"
  token: ""
  secret: ""
caasp:
  projectRules:
    moveCards:
      - action: labeled
        match: Blocked
        from: 
          - 1234566
          - 6654321
        destination: 7654321
      - action: labeled
        match: needinfo
        from: 
          - 1234566
          - 6654321
        destination: 7654321
      - action: labeled
        match: BugSquad
        from:
        destination: 6654321
      - action: unlabeled
        match: needinfo
        from:
          - 7654321
        destination: 1234566
      - action: assigned
        match:
        from:
          - 6654321
          - 7654321
        destination: 1234566
      - action: unassigned
        match:
        from:
          - 1234566
        destination: 7654321
      - action: closed
        match:
        from:
        destination: 98766543
```

In the server section we have configuration for the HTTP server like port and timeouts.

In the github section we have the configuration for the GitHub API.

The `CaaSP` section is specific to the caasp team, we maintained it and put what we want there, in this case are rules to manage our project in GitHub.

# Examples / User Stories
[examples]: #examples

## CaaSP Use Cases

### Action: Issue gets assigned

When this happens a Webhook is fired, when we get this, we move the respective card in a Project board to the `In Progress` column.

### Action: Issue gets labeled `neddinfo`

When this happens a Webhook is fired, when we get this, we move the respective card in a Project board from the `In Progress` column to the `Blocked` column.

# Drawbacks
[drawbacks]: #drawbacks

Why should we **not** do this?

- Security, we have to be sure, wi implement the correct security to avoid compromising any private data.

# Alternatives
[alternatives]: #alternatives

- Keep using project `Webhooks` but this will limit the reach of this.
  - This forces each team to produce their own tool.
  - Multi-repository project management becomes less manageable, for the Caasp team.

# Unresolved questions
[unresolved]: #unresolved-questions

- Have there been other attempts to do this?
