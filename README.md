# parking
Parking Lot

Binaries for common platforms can be downloaded from 
https://drive.google.com/drive/folders/1BY0E99toQb8Di_St6bbp_tk94f3ZmKJp?usp=sharing

Executing Binary:

Illustration for macOS
1. Download `parkingMACOS` from above source
2. Execute `./parkingMACOS -input <input_file_path>` (Eg : ./parkingMACOS -input input.txt)
3. If no filename is provided, it tries to open `input.txt` in same directory as the executable file.
(Eg: ./parkingMACOS) will try to read instructions from file `input.txt` in the same directory if available.

In case you need binary for platform not included in above, Please get in touch with me with your operating system and architecture information
or else binary can also be generated using below steps:

Generating Binary:
1. Clone this repo
2. `env GOOS=target-OS GOARCH=target-architecture go build`(replace `target-os` and `target-architecture` of your machine. Refer Table 1 below for additional information)
3. You might have to `go get` certain packages before running step 2, The error output of step 2 is self explanatory.
