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



Further Scaling options:

1. Implementing cache
The Project uses In Memory Array for storage of parking lot information.
For every search we iterate over the whole array. Should not be a problem if no of slots in a parking lot are less than 1000.
1st level optimisation to improve the searches would be to maintain a `cache` similar to a DB Index in the form of lets say a `MAP` with vehicle number as key and slot id as value. This will save us from full array scan and add the overhead of updating the map only when a vehicle enters or leaves a spot, which should not be a problem as we have a trigger point for this action.
MAP[vehicle_number] = slotID

```
lpmap := make(map[string]int)
//`KA-01-aa-1234` is parked at slot ID 2

lpmap["KA-01-AA-1234"] = 2

//`KA-01-AA-1234` leaving slot ID 2
delete(lpmap,"KA-01-AA-1234")

```

We'll also need to add another MAP with age as key and map of vehicle number as the value.
Adding second map in the above code
```
lpmap := make(map[string]int)
agemap := make(map[int]map[string]int)
//`KA-01-aa-1234` is parked at slot ID 2 driver age 21

lpmap["KA-01-AA-1234"] = 2
car := make(map[string]int)
car["KA-01-AA-1234"] = 2
agemap[21] = car

//`KA-01-AA-1234` leaving slot ID 2
delete(lpmap,"KA-01-AA-1234")
carsWithAge := agemap[21]
delete(carsWithAge,"KA-01-AA-1234")
agemap[21] = carsWithAge

//search for slot of car with licensePlate "KA-01-AA-1234"
answer := lpmap["KA-01-AA-1234"] 


//search for cars with drivers of age 21
result := ""
cars21 := agemap[21]
for licensePlate,_ := range cars21 {
  result += licensePlate
}

```


We can opt for this optimisation if our system is `search` heavy and not cars entering or leaving heavy.
In case the number of operations of cars entering and leaving the parking lot are more than searches this additional read and write overhead of maintaining the cace can be avoided.
