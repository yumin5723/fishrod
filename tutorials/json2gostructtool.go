package tutorials

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tidwall/gjson"
)

var (
	recursiveFlag = flag.Bool("r", false, "if r is set , output struct will be in closure format. Append to the end when use simplified format")
	jsonContent   = ""
)

type Task struct {
	name    string
	content map[string]interface{}
}

var (
	taskList []Task = make([]Task, 0)
)

func parseArgs(data string) {
	jsonContent = data
}

func handleParse(prefix string) {
	m, ok := gjson.Parse(jsonContent).Value().(map[string]interface{})
	if !ok {
		fmt.Println("json file parse error , please check the file format")
		os.Exit(0)
	}

	if prefix == "" {
		prefix = "Go"
	}
	taskList = append(taskList, Task{prefix, m})
}

func handleGoGenerate(prefix string) string {
	var structBuffer bytes.Buffer
	for {
		if len(taskList) == 0 {
			break
		}
		task := taskList[0]
		taskList = taskList[1:]
		content, _ := HandleTask(task, prefix)
		structBuffer.WriteString(content)
	}
	return structBuffer.String()
}

func HandleTask(task Task, prefix string) (res string, err error) {
	strcutName := task.name
	if prefix != "" {
		strcutName = strings.ToUpper(prefix[:1]) + prefix[1:] + task.name
	}
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf(`type %s struct {
`, strcutName))
	defer func() {
		buffer.WriteString("}\n")
		res = buffer.String()
	}()
	for key, val := range task.content {
		line, err := getStructLineString(key, val, task)
		if err != nil {
			panic(err.Error())
		}
		buffer.WriteString(line)
	}
	return buffer.String(), nil
}

func getStructLineString(key string, val interface{}, task Task) (line string, err error) {
	// write variable name
	oldKey := key
	var lineBuffer bytes.Buffer

	switch len(key) {
	case 0:
		return "", nil
	case 1:
		key = strings.ToUpper(key)
	default:
		key = strings.ToUpper(key[0:1]) + key[1:]

	}

	lineBuffer.WriteString(fmt.Sprintf("\t%s\t", key))
	// handle json value `null`
	if val == nil {
		val = struct{}{}
	}

	var (
		tp         = reflect.TypeOf(val)
		typeStr    string
		parseValue = task.content[oldKey]
	)

	//recursive handle function
	type TaskType string
	const (
		TaskSlice TaskType = "slice"
		TaskMap   TaskType = "map"
	)

	handleRecursiveTask := func(task Task, tp TaskType) (res string, err error) {
		bf := bytes.Buffer{}
		bf.WriteString(fmt.Sprintf("\t%s\t", key))
		switch tp {
		case TaskSlice:
			bf.WriteString("[]struct{\n")
		case TaskMap:
			bf.WriteString("struct{\n")
		default:
			panic("wrong task type ")
		}
		defer func() {
			bf.WriteString("\t}\t")
			bf.WriteString(fmt.Sprintf("`json:\"%s\"`\n", oldKey))
			res = bf.String()
		}()
		for key, val := range task.content {
			line, err := getStructLineString(key, val, task)
			if err != nil {
				panic(err)
			}
			bf.WriteString("\t" + line)
		}

		return bf.String(), nil
	}

	for tp.Kind() == reflect.Slice {
		mps, ok := parseValue.([]interface{})
		if !ok {
			return "", fmt.Errorf("wrong value type : %s ", val)
		}
		if len(mps) == 0 {
			break
		}
		fmt.Println(mps)
		// 判断[]内类型是否相同，不相同则为[]interface{}
		baseType := reflect.TypeOf(mps[0])
		isStandardArray := true
		for _, mp := range mps {
			if reflect.TypeOf(mp) != baseType {
				isStandardArray = false
				break
			}
		}
		// 数组内类型不一样，go文件中生成[]interface{}
		typeStr += "[]"
		if !isStandardArray {
			tp = reflect.TypeOf(struct{}{})
		} else if baseType.Kind() == reflect.Slice {
			parseValue = mps[0]
			continue
		} else if baseType.Kind() == reflect.Map {
			name := key + "Item"
			typeStr += name

			unionMap := getUnionFieldMap(mps)
			if *recursiveFlag {
				return handleRecursiveTask(Task{name, unionMap}, TaskSlice)
			}

			taskList = append(taskList, Task{name, unionMap})
			goto Output
		} else {
			parseValue = mps[0]
			tp = reflect.TypeOf(parseValue)
		}

	}

	// 输出类型
	switch tp.Kind() {
	case reflect.Bool:
		typeStr += "bool"
	case reflect.String:
		typeStr += "string"
	case reflect.Int:
		typeStr += "int"
	case reflect.Float32:
		typeStr += "float32"
	case reflect.Float64:
		typeStr += "float64"
	case reflect.Struct: // 支持null值
		typeStr += "interface{}"
	case reflect.Slice: // 支持null值
		typeStr += "[]interface{}"
	case reflect.Map:
		typeStr += key + "Item"
		mp, ok := parseValue.(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("wrong value type : %s ", val)
		}
		if *recursiveFlag {
			return handleRecursiveTask(Task{typeStr, mp}, TaskMap)
		} else {
			taskList = append(taskList, Task{typeStr, mp})
		}

	default:
		return "", fmt.Errorf("wrong value type : %s ", val)
	}

Output:
	lineBuffer.WriteString(typeStr + "\t")

	// write tag
	tagStr := fmt.Sprintf("`json:\"%s\"`\n", oldKey)
	lineBuffer.WriteString(tagStr)

	return lineBuffer.String(), nil
}

func getUnionFieldMap(mps []interface{}) (unionMap map[string]interface{}) {
	unionMap = make(map[string]interface{})
	for _, v := range mps {
		v, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		for key, field := range v {
			if _, ok := unionMap[key]; ok {
				continue
			}
			unionMap[key] = field
		}

	}
	return
}

func json2gostructtoolScreen(_ fyne.Window) fyne.CanvasObject {
	currentArea := widget.NewMultiLineEntry()
	currentArea.Wrapping = fyne.TextWrapWord

	targetArea := widget.NewMultiLineEntry()
	targetArea.Wrapping = fyne.TextWrapWord

	prefix := widget.NewEntry()
	prefix.PlaceHolder = "前缀"

	buttonBox := container.NewVBox(prefix, widget.NewButton("转换", func() {
		parseArgs(currentArea.Text)
		handleParse(prefix.Text)
		// fmt.Println(err.Error())
		targetArea.SetText(handleGoGenerate(prefix.Text))
	}),
	)

	// content := container.NewVBox(
	// 	currentArea, buttonBox, targetArea,
	// )

	return container.NewGridWithRows(1, currentArea, buttonBox, targetArea)
}
