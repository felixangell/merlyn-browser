# merlyn
toy web browser!

	css/
		css parser - aiming to only just be CSS2 compliant
	dom/
		document object model thing
	frontend/
		source code for the frontend
		of the engine, this is the part
		that shows the window and handles
		the user interface-y stuff
	html/
		html parser - compliant with HTML 1999 standards if im lucky

## building and installing
I'm using Mac Sierra:

	brew install pkg-config
	brew install sdl2
	brew install sdl2_{ttf,img}

	go get github.com/felixangell/merlyn

You may have to get the strife library too if it doesn't automagically download somehow:

	go get github.com/felixangell/strife

