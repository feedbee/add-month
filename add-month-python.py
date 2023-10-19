from datetime import datetime, timedelta
from dateutil.relativedelta import relativedelta

# Set a specific date to October 31, 2023
today = datetime.now()
# Add one month to the specific date
future_date = today + relativedelta(months=1)

print("Specific Date:", today.strftime('%Y-%m-%d'))
print("Date One Month from Now:", future_date.strftime('%Y-%m-%d'))

print("")

# Set a specific date to October 31, 2023
specific_date = datetime(2023, 10, 31)
# Add one month to the specific date
future_date2 = specific_date + relativedelta(months=1)

print("Specific Date:", specific_date.strftime('%Y-%m-%d'))
print("Date One Month from Now:", future_date2.strftime('%Y-%m-%d'))
