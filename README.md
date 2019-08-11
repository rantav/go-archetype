# go-archetype

go-archetype is a tool for creating project archetypes.

If you or your company create new projects and you want all these projects to preserve a similar structure or simply create a good starting point for other develoeprs, use go-archetype to define project archtypes (templates).

## Concept

Two roles are involved:

* Archetype creator (typically a senior developer/architect)
* Developer creating a new project

### Archetype creator

As the archetype creator you create a template (aka blueprint) project. This project is fully functional, valid code written in Golang, JavaScript, Java, Python etc. (go-archetype is agnostic)
And then you define a set of transformaitons. E.g. ask the user for the project name and replace here, here and here.

### Developer using the archetype

As a developer creating a new project based on the archetype, you run `go-archetype` with a few parameters, you answer a few questions (such as what if you new project's name, these questions are defined by the archetype creator) and your new code is generated.

### Summary

Here are the steps:

1. Architect creates an actual working project. This project is used as a blueprint or an archetype. It's a barebones project, no templating involved.
1. Architect defines a set of transformations in transformations.yml. These transformations are used for simple search-and-replace in the blueprint. (and there are more compicated types, see below)
1. Developer runs go-archetype to create a new project based on the blueprint.

## Usage

    go-archetype --transformaitons=transformations.yml --source=. --destination=/path/to/your/new/project

## Example transformaitons.yml file

    TODO

## Requesting for user input

    TODO

## Operations and debugging

The view detailed logs, run with `LOG_LEVEL=debug`
