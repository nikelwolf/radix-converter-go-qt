package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"

	"github.com/nikelwolf/radix-converter-go-qt/converter"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	quickcontrols2.QQuickStyle_SetStyle("Material")

	engine := qml.NewQQmlApplicationEngine(nil)

	rc := converter.NewRadixConverter(nil)
	rc.QmlRegisterType()
	engine.RootContext().SetContextProperty("_converter", rc)

	engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	gui.QGuiApplication_Exec()
}
