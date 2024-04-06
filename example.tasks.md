---
TODO:
  - personal:single task
---

# personal

## single task

- compiler - .tasks.md file to json (`$ stask compile <file> -o <output file>`) (output defaults to stdout)
  json fields: current, todo, tasks.
  tasks - structured tasks object, with keys as category, each category has keys as subjects, and those are arrays of objects of task and description.
  todo - in order of to do, array of objects with category, subject, task, description.
  current - first item of todos.
  there will be data duplication but i think its not a big deal.
- lsp - autocomplete and diagnostics
  autocomplete for authoring TODO section based of existing tasks.
  diagnostics for TODO and TASKS sections to follow correct syntax.
- vscode extension
  lsp client
  syntax highlighting
  customization options:
  tasks file location (can be set for workspace with .vscode/setting.json)
  where to display the task (eg. status bar)
  highlight color for the status bar task (from https://code.visualstudio.com/api/references/theme-color)
  template for displaying the task, eg: `{category}/{subject}: {task}` or just `{task}` (default)
- syntax highlighting notes
  highlight categories/subjects/tasks in different colors. same highlights apply in the TODO section.
  highlight the section in TASKS that is currently active (based on the TODO section / sequential order).
  maybe the other section could maintain the colors but be less opaque.
  hover on item in TODOS should background-highlight the group that it references

## nvim config

- add git sync statusbar item next to branch
  the item will show the number of commits behind remote
  `git fetch` we be ran on nvim open and periodically to check
