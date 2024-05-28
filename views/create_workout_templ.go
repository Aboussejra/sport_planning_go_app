// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CreateWorkout() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div><h2 class=\"text-2xl font-bold text-center mb-8\">Create Workout</h2><form action=\"/create-workout\" method=\"post\" id=\"workout-form\" class=\"mb-8\"><div class=\"mb-4\"><label for=\"name\" class=\"block text-gray-700\">Workout Name:</label> <input type=\"text\" id=\"name\" name=\"name\" class=\"form-input mt-1 block w-full\" required></div><div class=\"mb-4\"><label for=\"dayOfWeek\" class=\"block text-gray-700\">Day of the Week:</label> <select id=\"dayOfWeek\" name=\"dayOfWeek\" class=\"form-select mt-1 block w-full\" required><option value=\"Monday\">Monday</option> <option value=\"Tuesday\">Tuesday</option> <option value=\"Wednesday\">Wednesday</option> <option value=\"Thursday\">Thursday</option> <option value=\"Friday\">Friday</option> <option value=\"Saturday\">Saturday</option> <option value=\"Sunday\">Sunday</option></select></div><button type=\"submit\" class=\"bg-blue-700 hover:bg-blue-900 text-white font-bold py-2 px-4 rounded\">Create Workout</button></form></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
