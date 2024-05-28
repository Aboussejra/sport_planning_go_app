// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.696
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "time"

func Index(currentDate time.Time) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>My Workout App</title><link href=\"https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css\" rel=\"stylesheet\"><script src=\"https://unpkg.com/htmx.org@1.5.0\"></script></head><body class=\"bg-gray-100\"><div class=\"container mx-auto py-8\"><h1 class=\"text-3xl font-bold text-center mb-8\">Workout Scheduling App</h1><!-- Button container --><div class=\"flex justify-center mb-8\"><button id=\"create_workout_button\" class=\"bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded mr-4\" onclick=\"toggleDiv(&#39;create_workout&#39;)\">Create Workout</button> <button id=\"add_exercise_button\" class=\"bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded\" onclick=\"toggleDiv(&#39;add_exercises&#39;)\">Add Exercises</button><!-- Add more buttons here as needed --></div><!-- Divs to be displayed conditionally --><div id=\"create_workout\" class=\"toggle-content hidden\"><p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CreateWorkout().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div><div id=\"add_exercises\" class=\"toggle-content hidden\"><p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = AddExercises().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div></div><script>\n        function toggleDiv(divId) {\n            var buttonLists = ['create_workout', 'add_exercises']; // Add more div IDs as needed\n\n            buttonLists.forEach(function(buttonId) {\n                var div = document.getElementById(buttonId);\n                if (divId === buttonId) {\n                    div.classList.remove(\"hidden\");\n                } else {\n                    div.classList.add(\"hidden\");\n                }\n            });\n        }\n    </script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
