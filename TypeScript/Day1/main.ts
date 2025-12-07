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
    let increment: number = 0;
    
    increment = Math.trunc(lock + steps / 100);
    if(direction === 'R'){
      lock += steps;
      lock %= 100;
    }else if(direction === 'L'){
      lock -= steps;
      lock %= 100;  
    }

    
  }
  return count;
}

console.log(decodeMessage(fileContent));