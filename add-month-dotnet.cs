using System;
					
public class Program
{
	public static void Main()
	{
		// Get the current date and time
		DateTime currentDate = DateTime.Now;
		// Add one month to the current date
		DateTime futureDate = currentDate.AddMonths(-1);

		Console.WriteLine("Current Date: " + currentDate);
		Console.WriteLine("Date One Month from Now: " + futureDate);
		
		Console.WriteLine("");
		
		// Set a specific date to October 31, 2023
		DateTime specificDate = new DateTime(2023, 10, 31);
		// Add one month to the specific date
		DateTime futureDate2 = specificDate.AddMonths(-1);
		
		Console.WriteLine("Specific Date: " + specificDate);
		Console.WriteLine("Date One Month from the Specific Date: " + futureDate2);
	}
}
