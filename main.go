package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"os"
	"strings"
	"time"

	"./dreamcolor"
	"github.com/EdlinOrg/prominentcolor"
	"github.com/kbinani/screenshot"
	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
)

var finishTask = make(chan struct{})

func onStateChanged(device gatt.Device, state gatt.State) {

	log.Println("[onStateChanged]", state)

	switch state {
	case gatt.StatePoweredOn:
		device.Scan([]gatt.UUID{}, false)
		return
	default:
		device.StopScanning()
	}
}

func onPeripheralDiscovered(peripheral gatt.Peripheral, a *gatt.Advertisement, rssi int) {

	id := strings.ToUpper(flag.Args()[0])
	fmt.Printf("\nPeripheral: %s (%s)\n\n", peripheral.Name(), peripheral.ID())

	if strings.ToUpper(peripheral.ID()) != id {
		return
	}

	device := peripheral.Device()
	device.StopScanning()
	device.Connect(peripheral)
}

func findServiceByUUID(peripheral gatt.Peripheral, uuid string) *gatt.Service {

	controlServiceUUID, _ := gatt.ParseUUID(uuid)
	services, _ := peripheral.DiscoverServices([]gatt.UUID{controlServiceUUID})

	for _, service := range services {
		if service.UUID().Equal(controlServiceUUID) {
			return service
		}
	}

	return nil
}

func findCharacteristicByType(peripheral gatt.Peripheral, service *gatt.Service, property gatt.Property) *gatt.Characteristic {

	characteristics, _ := peripheral.DiscoverCharacteristics([]gatt.UUID{}, service)

	for _, characteristic := range characteristics {
		if (characteristic.Properties() & property) != 0 {
			return characteristic
		}
	}

	return nil
}

func write(peripheral gatt.Peripheral, characteristic *gatt.Characteristic, command []byte) error {
	return peripheral.WriteCharacteristic(characteristic, command, true)
}

func onPeripheralConnected(peripheral gatt.Peripheral, exception error) {

	defer peripheral.Device().CancelConnection(peripheral)

	if exception := peripheral.SetMTU(500); exception != nil {
		fmt.Printf("[onPeripheralConnected] Failed to set MTU, exception: %s\n", exception)
	}

	// services, exception := peripheral.DiscoverServices(nil)

	// if exception != nil {
	// 	fmt.Printf("[onPeripheralConnected] Failed to discover services, exception: %s\n", exception)
	// 	return
	// }

	notificationService := findServiceByUUID(peripheral, "000102030405060708090a0b0c0d1910")
	writeCharacteristic := findCharacteristicByType(peripheral, notificationService, gatt.CharWriteNR)

	//setAmbilight(peripheral, writeCharacteristic)

	write(peripheral, writeCharacteristic, dreamcolor.SetColorAlternate(dreamcolor.ColorCommand{
		ColorA: dreamcolor.RgbColor{Red: 255, Green: 0, Blue: 0},
		ColorB: dreamcolor.RgbColor{Red: 0, Green: 255, Blue: 255},
		Toggle: true,
	}))

	// for _, service := range services {

	// 	serviceName := service.Name()
	// 	serviceInformation := "Service: " + service.UUID().String()

	// 	if len(serviceName) > 0 {
	// 		serviceInformation += " (" + serviceName + ")"
	// 	}

	// 	fmt.Println(serviceInformation)
	// 	characteristics, exception := peripheral.DiscoverCharacteristics(nil, service)

	// 	if exception != nil {
	// 		log.Printf("Failed to discover characteristics, exception: %s\n", exception)
	// 		continue
	// 	}

	// 	for _, characteristic := range characteristics {

	// 		characteristicInformation := "\tCharacteristic  " + characteristic.UUID().String()

	// 		if len(characteristic.Name()) > 0 {
	// 			characteristicInformation += " (" + characteristic.Name() + ")"
	// 		}

	// 		characteristicInformation += "\n\t\tProperties    " + characteristic.Properties().String()
	// 		fmt.Println(characteristicInformation)

	// 		if (characteristic.Properties() & gatt.CharRead) != 0 {

	// 			b, exception := peripheral.ReadCharacteristic(characteristic)

	// 			if exception != nil {
	// 				log.Printf("Failed to read characteristic, exception: %s\n", exception)
	// 				continue
	// 			}

	// 			fmt.Printf("\t\tValue         %x | %q\n", b, b)
	// 		}

	// 		descriptors, exception := peripheral.DiscoverDescriptors(nil, characteristic)

	// 		if exception != nil {
	// 			log.Printf("Failed to discover descriptors, exception: %s\n", exception)
	// 			continue
	// 		}

	// 		for _, d := range descriptors {
	// 			msg := "  Descriptor      " + d.UUID().String()
	// 			if len(d.Name()) > 0 {
	// 				msg += " (" + d.Name() + ")"
	// 			}
	// 			fmt.Println(msg)

	// 			// Read descriptor (could fail, if it's not readable)
	// 			b, err := peripheral.ReadDescriptor(d)
	// 			if err != nil {
	// 				log.Printf("Failed to read descriptor, err: %s\n", err)
	// 				continue
	// 			}
	// 			fmt.Printf("    value         %x | %q\n", b, b)
	// 		}

	// 		// Subscribe the characteristic, if possible.
	// 		if (characteristic.Properties() & (gatt.CharNotify | gatt.CharIndicate)) != 0 {
	// 			f := func(c *gatt.Characteristic, b []byte, err error) {
	// 				fmt.Printf("notified: % X | %q\n", b, b)
	// 			}
	// 			if err := peripheral.SetNotifyValue(characteristic, f); err != nil {
	// 				fmt.Printf("Failed to subscribe characteristic, err: %s\n", err)
	// 				continue
	// 			}
	// 		}

	// 	}
	// 	fmt.Println()
	// }
}

func onPeripheralDisconnected(peripheral gatt.Peripheral, exception error) {
	fmt.Println("onPeripheralDisconnected")
	//peripheral.Device().Connect(peripheral)
	close(finishTask)
}

func main() {

	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatalf("usage: %s [options] peripheral-id\n", os.Args[0])
	}

	device, exception := gatt.NewDevice(option.DefaultClientOptions...)

	if exception != nil {
		log.Fatalf("Failed to open device, exception: %s\n", exception)
	}

	device.Handle(
		gatt.PeripheralDiscovered(onPeripheralDiscovered),
		gatt.PeripheralConnected(onPeripheralConnected),
		gatt.PeripheralDisconnected(onPeripheralDisconnected),
	)

	device.Init(onStateChanged)
	<-finishTask
	fmt.Println("mainDone")
}

func processBatch(k int, flags int, screenshot image.Image) prominentcolor.ColorRGB {

	dimensions := uint(prominentcolor.DefaultSize)
	masks := prominentcolor.GetDefaultMasks()
	result, _ := prominentcolor.KmeansWithAll(k, screenshot, flags, dimensions, masks)

	return result[0].Color
}

func setAmbilight(peripheral gatt.Peripheral, characteristic *gatt.Characteristic) {

	var previous dreamcolor.RgbColor

	for {
		screnshot, _ := screenshot.CaptureRect(screenshot.GetDisplayBounds(0))
		prominent := processBatch(1, prominentcolor.ArgumentNoCropping, screnshot)

		current := dreamcolor.RgbColor{
			Red:   prominent.R,
			Green: prominent.G,
			Blue:  prominent.B,
		}

		if current != previous {
			write(peripheral, characteristic, dreamcolor.SetColor(current))
		}

		time.Sleep(100 * time.Millisecond)
	}
}
