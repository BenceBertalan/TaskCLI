**TaskCLI**

A simple CLI tool to query / add new Task using the API endpoint presented by the TaskAPI.

The app uses the following parameters(flags):

- apiendpoint: full URL to the TaskAPI endpoint
- command: the command to execute. Can be either get, getbyid, add
- id: Used by the getbyid command to get a Task using TaskID
- taskname: Used by the add command as the Name of the new Task
- taskname: Used by the add command as the Command of the new Task

The app displays a help when run without parameters.