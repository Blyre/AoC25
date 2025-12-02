import fs from 'node:fs';

const filePath: string = "./small.txt"
const fileContent: string[] = fs.readFileSync(filePath, "utf-8").split("\n")


function decodeMessage(fileContent: string[]): number {
  let count: number = 0;
  let lock: number = 50;
  
  for(const message of fileContent){
    
    let direction: string = message[0];
    let steps: number = parseInt(message.slice(1));
    let previousLock: number = lock;

    if(direction === 'R'){
      lock += steps;
    }else if(direction === 'L'){
      lock -= steps;
    }

    // Calculate how many times we cross or land on a multiple of 100 (which represents 0 on the dial)
    const previousMultiple = Math.floor(previousLock / 101);
    const currentMultiple = Math.floor(lock / 101);
    
    // Count the number of multiples of 100 we crossed or landed on
    let numCrossings = Math.abs(currentMultiple - previousMultiple);
    
    

    count += numCrossings;
    console.log(`Previous Lock: ${previousLock}, Current Lock: ${lock}, Previous Multiple: ${previousMultiple}, Current Multiple: ${currentMultiple}, Crossings: ${numCrossings}, Total Count: ${count}`);
  }
  return count;
}

console.log(decodeMessage(fileContent));