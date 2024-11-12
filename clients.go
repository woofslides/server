package main

type Connection struct {
	Viewer     string
	Controller string
	Show       *Show
}

// map[controller address]viewer address
var pending map[string]string

// map[controller address]point to a Connection
var connections map[string]*Connection

// true or false is returned to tell the consuming function whether or not the connection is accepted
func addConnection(viewer string, controller string) bool {
	pendingViewer, ok := pending[controller]

	if ok {
		if pendingViewer != viewer {
			return false
		}
	} else {
		pendingViewer = viewer
	}

	connections[controller] = &Connection{Viewer: viewer, Controller: controller}
	delete(pending, controller)

	return true

}

// NEED TO DO SOMETHING IF THE CONNECTION MUST BE PENDING

// true or false is returned to tell the consuming function whether or not the connection exists
func addShow(controller string, show *Show) bool {
	connection, ok := connections[controller]

	if !ok {
		return false
	}

	connection.Show = show

	return true
}
