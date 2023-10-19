require 'date'

# Get the current date
current_date = Date.today

# Add one month to the current date
future_date = current_date >> 1

# Print the current date and the date one month from now
puts "Current Date: #{current_date}"
puts "Date One Month from Now: #{future_date}"

puts ""

# Set a specific date to October 31, 2023
specific_date = Date.new(2023, 10, 31)
# Add one month to the current date
future_date2 = specific_date >> 1

puts "Specific Date: #{specific_date}"
puts "Date One Month from the Specific Date: #{future_date2}"
