package handlers

import (
	"encoding/json"
	"log"

	// "tkoh_oms/models"

	errorHandler "tkoh_oms/errors"
	"tkoh_oms/models/systemStatus"
	ws_model "tkoh_oms/models/websocket"
	"tkoh_oms/models/wsTest"
	"tkoh_oms/websocket"

	"github.com/gofiber/fiber/v2"
)

const AW2_RESPONSE string = `{
	"messageCode": "LOCATION_UPDATE",
	"userId": 1,
	"dutyLocationId" : 1,
	"dutyLocationName" : "5/F DSC"
}`

// @Summary		Test AW2 websocket response.
// @Description	Get the response of AW2 (Server notify the user which location selected).
// @Tags			Test
// @Accept			json
// @Param param body wsTest.ReportDutyLocationUpdateResponse true "AW2 Parameters"
// @Produce		plain
// @Success		200	"OK"
// @Router			/testAW2 [post]
func HandleTestAW2(c *fiber.Ctx) error {

	// mqtt.PublishMqtt("direct/publish", []byte("packet scheduled message"))
	var response wsTest.ReportDutyLocationUpdateResponse
	if err := c.BodyParser(&response); errorHandler.CheckError(err, "Invalid/missing input: ") {
		err := json.Unmarshal([]byte(AW2_RESPONSE), &response)
		errorHandler.CheckError(err, "translate string to json in wsTest")
	}
	websocket.SendBoardcastMessage(response)
	return c.SendString("OK")
}

const OW1_RESPONSE string = `{
    "messageCode": "ORDER_STATUS",
    "scheduleId": 1,
    "orderId": 1,
    "robotId":["AMR1"],
    "payloadId": "CART1",
    "orderStatus": "PROCESSING",
    "processingStatus": "ARRIVED_START_LOCATION"
}`

// @Summary		Test OW1 websocket response.
// @Description	Get the response of OW1 (Server notify any of created order status changed).
// @Tags			Test
// @Accept			json
// @Param param body wsTest.ReportOrderStatusUpdateResponse true "OW1 Parameters"
// @Produce		plain
// @Success		200	"OK"
// @Router			/testOW1 [post]
func HandleTestOW1(c *fiber.Ctx) error {
	// get the processingStatus from the request body
	// processingStatus := c.Query("processingStatus")
	var response wsTest.ReportOrderStatusUpdateResponse
	if err := c.BodyParser(&response); errorHandler.CheckError(err, "Invalid/missing input: ") {
		err := json.Unmarshal([]byte(OW1_RESPONSE), &response)
		errorHandler.CheckError(err, "translate string to json in wsTest")
	}
	// err := json.Unmarshal([]byte(OW1_RESPONSE), &response)
	// fmt.Println(response)
	// response.ProcessingStatus = processingStatus
	// errorHandler.CheckError(err, "translate string to json in wsTest")
	websocket.SendBoardcastMessage(response)
	return c.SendString("OK")
}

const MW1_RESPONSE string = `{
    "messageCode": "ROBOT_STATUS",
    "robotList": [
        {
            "robotId": "VIRTUAL_AMR1",
            "robotState": "IDLE",
            "robotStatus": [
                "STOP",
                "OPERATION/LOADED"
            ],
            "zone": "LG/F",
            "robotPosition": [
                -2.85632,
                0.574238,
                0
            ],
            "robotOritentation": null,
            "robotCoordination": null,
            "batteryLevel": -1,
            "lastReportTime": "20240102031700"
        }
    ]
}`

// @Summary		Test MW1 websocket response.
// @Description	Get the response of MW1 (Server report robot status and location (every 1s) ).
// @Tags			Test
// @Accept			json
// @Param param body ws_model.WebsocketUpdateRobotStatusResponse true "MW1 Parameters"
// @Produce		plain
// @Success		200	"OK"
// @Router			/testMW1 [post]
func HandleTestMW1(c *fiber.Ctx) error {
	log.Print("HandleTestMW1")
	response := ws_model.WebsocketUpdateRobotStatusResponse{}
	// var response wsTest.ReportRobotStatusLocationResponse
	if err := c.BodyParser(&response); errorHandler.CheckError(err, "Invalid/missing input: ") {
		err := json.Unmarshal([]byte(MW1_RESPONSE), &response)
		errorHandler.CheckError(err, "translate string to json in wsTest")
	}
	// err := json.Unmarshal([]byte(MW1_RESPONSE), &response)
	// errorHandler.CheckError(err, "translate string to json in wsTest")
	websocket.SendBoardcastMessage(response)
	return c.SendString("OK")
}

const SW1_RESPONSE string = `{
    "messageCode": "SYSTEM_STATUS",
    "systemState": "STOPPED",
    "systemStatus": ["LIFT_ALARM"]
}`

// @Summary		Test SW1 websocket response.
// @Description	Get the response of SW1 (Server report robot status and location (every 1s) ).
// @Tags			Test
// @Accept			json
// @Param param body systemStatus.SystemStatusResponse true "SW1 Parameters"
// @Produce		plain
// @Success		200	"OK"
// @Router			/testSW1 [post]
func HandleTestSW1(c *fiber.Ctx) error {
	var response systemStatus.SystemStatusResponse
	if err := c.BodyParser(&response); errorHandler.CheckError(err, "Invalid/missing input: ") {
		err := json.Unmarshal([]byte(SW1_RESPONSE), &response)
		errorHandler.CheckError(err, "translate string to json in wsTest")
	}
	websocket.SendBoardcastMessage(response)
	return c.SendString("OK")
}

// const WW1_RESPONSE string = `{
//     "messageCode": "SYSTEM_STATUS",
//     "systemState": "STOPPED",
//     "systemStatus": ["LIFT_ALARM"]
// }`

// // @Summary		Test WW1 websocket response.
// // @Description	Get the response of WW1 (Server report order status when update occurs).
// // @Tags			Test
// // @Accept			json
// // @Param param body systemStatus.SystemStatusResponse true "WW1 Parameters"
// // @Produce		plain
// // @Success		200	"OK"
// // @Router			/testWW1 [post]
// func HandleTestWW1(c *fiber.Ctx) error {
// 	var response systemStatus.SystemStatusResponse
// 	if err := c.BodyParser(&response); errorHandler.CheckError(err, "Invalid/missing input: ") {
// 		err := json.Unmarshal([]byte(SW1_RESPONSE), &response)
// 		errorHandler.CheckError(err, "translate string to json in wsTest")
// 	}
// 	websocket.SendBoardcastMessage(response)
// 	return c.SendString("OK")
// }
