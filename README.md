
# Chip-8 Emulator
A chip-8 emulator made in golang with the sdl2 graphics libary

### Tested with: 
- BC-Test Rom By BestCoder
- IBM-Logo Rom 
- chip8-test-rom By corax89
- Pong

### Installation
Make sure to have go and sdl2 installed, refer to [here](https://github.com/veandco/go-sdl2#requirements) for sdl-2 library instructions.

Clone repo
```bash
git clone https://github.com/shivkar2n/CHIP-8-Emulator.git
```

### Build and run emulator
In repo directory,
```bash
cd src && go build -o ../build
cd ../build
./Chip8-Emulator ../roms/Pong.ch8
```

### References
I followed [this](https://tobiasvl.github.io/blog/write-a-chip-8-emulator/) writeup, as well as this chip-8 [reference guide](http://devernay.free.fr/hacks/chip8/C8TECH10.HTM), credits go to their respective authors.