// Get the current date
var currentDate = new Date();
// Add 1 month
var futureDate = new Date(currentDate);
futureDate.setMonth(currentDate.getMonth() + 1);

console.log("Today:", currentDate)
console.log("Date One Month from Now:", futureDate)

console.log("")

// Set a specific date to October 31, 2023
var specificDate = new Date(2023, 9, 31); // Note: Months are zero-based, so 9 represents October
// Add 1 month
var futureDate2 = new Date(specificDate);
futureDate2.setMonth(specificDate.getMonth() + 1);

console.log("Specific Date:", specificDate)
console.log("Date One Month from the Specific Date:", futureDate2)
