package main

type Connection struct {
	Viewer     string
	Controller string
	Show       *Show
}

var pendingViewers map[string]string = make(map[string]string)     // map[controller address]viewer address
var pendingControllers map[string]string = make(map[string]string) // map[viewer address]controller address

var connections map[string]*Connection = make(map[string]*Connection) // map[controller address]point to a Connection

// -1: viewer/controller is already taken
//  0: added controller/viewer to waiting list
//  1: viewer and controller connection initialised

func viewerRequestController(viewer string, controller string) int {
	pendingController, ok := pendingControllers[viewer]

	if ok {
		if pendingController != controller {
			return -1
		}
	} else {
		pendingViewers[controller] = viewer
		return 0
	}

	connections[controller] = &Connection{Viewer: viewer, Controller: controller}
	delete(pendingControllers, viewer)

	return 1

}

func controllerRequestViewer(controller string, viewer string) int {
	pendingViewer, ok := pendingViewers[controller]

	if ok {
		if pendingViewer != viewer {
			return -1
		}
	} else {
		pendingControllers[viewer] = controller
		return 0
	}

	connections[controller] = &Connection{Viewer: viewer, Controller: controller}
	delete(pendingViewers, controller)

	return 1

}

// true or false is returned to tell the consuming function whether or not the connection exists
func addShow(controller string, show *Show) bool {
	connection, ok := connections[controller]

	if !ok {
		return false
	}

	show.Viewer = connection.Viewer
	show.Controller = connection.Controller
	connection.Show = show

	return true
}
