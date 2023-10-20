import java.time.LocalDate;

public class Main {
    public static void main(String[] args) {

        // Get the current date
        LocalDate currentDate = LocalDate.now();
        // Add 1 month
        LocalDate futureDate = currentDate.plusMonths(1);

        System.out.println("Current Date: " + currentDate);
        System.out.println("Date One Month from Now: " + futureDate);

        System.out.println("");
        
        // Set a specific date to January 31, 2024
        LocalDate specificDate = LocalDate.of(2024, 1, 31);
        // Add 1 month
        LocalDate futureDate2 = specificDate.plusMonths(1);

        System.out.println("Specific Date: " + specificDate);
        System.out.println("Date One Month from the Specific Date: " + futureDate2);
    }
}
