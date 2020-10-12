SRC_FILES := $(ls -I "main.go")

all: build
	@echo "Starting emulator, this might take some time..."
	@./tools/qemu-stm32 -machine stm32-f103c8 -kernel kernel.bin --nographic
	@make clean

flash:
	@tinygo flash -target=bluepill main.go

build:
	@echo "Building..."
	@tinygo build -target=bluepill -o kernel.bin $(SRC_FILES)

clean:
	@rm -f kernel.bin core DAC*