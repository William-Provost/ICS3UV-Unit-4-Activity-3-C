/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-30
 * @fileoverview This program finds average temperature, using arrays.
 */

// constants and variables
const temperature1: number[] = new Array(5); // Creates an array of 5 elements
let sum1: number = 0;

// get five temperatures from user
for (let counter = 0; counter < 5; counter++) {
  const userInput = prompt(`Enter temperature ${counter + 1}: `) || "0";
  const oneTemperature1 = parseFloat(userInput);
  temperature1[counter] = oneTemperature1;
}

// calculate sum using a for loop
for (let counter = 0; counter < temperature1.length; counter++) {
  sum1 += temperature1[counter];
}

// calculate average to two decimal places
// toFixed(2) returns a string, so we keep the const type as const
const average1 = (sum1 / temperature1.length).toFixed(2);

// display all temperatures using a for loop
console.log("\nThe temperatures are:");
for (let counter = 0; counter < temperature1.length; counter++) {
  console.log(`Temperature ${counter + 1}: ${temperature1[counter]}°C`);
}

// display the average
console.log(`\nThe average temperature is: ${average1}°C`);

console.log("\nDone.");
