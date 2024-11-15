package main

import "fmt"

func simulateExampleWorkflow() {
	result_for_controller := controllerRequestViewer("controlleraddress0000", "vieweraddress0000")
	fmt.Printf("Result from controllerRequestViewer 0: %d\n", result_for_controller)

	result_for_viewer := viewerRequestController("vieweraddress0000", "controlleraddress0000")
	fmt.Printf("Result from viewerRequestController 0: %d\n", result_for_viewer)

	result_for_new_viewer := viewerRequestController("vieweraddress1111", "controlleraddress1111")
	fmt.Printf("Result from viewerRequestController 1: %d\n", result_for_new_viewer)

	result_for_new_controller := controllerRequestViewer("controlleraddress1111", "vieweraddress1111")
	fmt.Printf("Result from controllerRequestViewer 1: %d\n", result_for_new_controller)

	fmt.Println(pendingControllers)
	fmt.Println(pendingViewers)
	fmt.Println(connections)
}
