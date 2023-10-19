<?php

// Get the current date
$today = (new DateTimeImmutable());
// Add 1 month
$nextMonth = $today->add(new DateInterval('P1M'));

echo "Current Date: " . $today->format('c') . "\n";
# Add one month to the specific date
echo "Date One Month from Now: " . $nextMonth->format('c') . "\n";

echo "\n";

# Set a specific date to October 31, 2023
$specificDate = (new DateTimeImmutable())->setDate(2023, 10, 31)->setTime(0, 0, 0);
# Add one month to the specific date
$nextMonth2 = $specificDate->add(new DateInterval('P1M'));

echo "Specifit Date: " . $specificDate->format('c') . "\n";
echo "Date One Month from the Specific Date: " . $nextMonth2->format('c') . "\n";
