const array = [4, 7, 2, 9, 5]
let sum = 0;

function findAverage(array) {
    for (let i = 0; i < array.length; i++){
        sum += array[i]
    }
    sum = sum / array.length
    return sum 
}

console.log(findAverage(array));
