all:
	go run main.go draw.go matrix.go line.go transform.go parse.go parametric.go shapes.go stack.go screen.go lighting.go
clean:
	rm *.png
	rm *.ppm
	rm parse/*.pyc
