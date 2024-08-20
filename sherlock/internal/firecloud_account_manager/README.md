# firecloud_account_manager

This package contains code to manage Firecloud accounts. What's unique about this code is that it isn't integrated with the rest of the codebase -- even parts that might make sense, like role_propagation or suitability_synchronization.

The context is that this code was written on a tight timeline so it could be documented and finalized before audit activities. We already had code that helped manage Firecloud accounts but it was in Jenkins, and there were a few significant problems with that (mainly security and reliability). Around the time that those issues began to boil over, Sherlock gained both production-approval and access to the BITS data warehouse, opening it up as another potential location for this code.

The benefit of this code being here is that it can piggyback off of the rest of Sherlock's security, reliability, and test structure.

This code does three things:
- If an account hasn't been logged in to for the first time and it's older than the configured grace period, the account is suspended.
- If an account hasn't been logged in to since the inactivity threshold, the account is suspended.
- If an account doesn't have a matching user in the BITS data warehouse, the account is suspended.

Whenever an account is suspended, notifications are sent to the security events channel(s). The code respects a list of accounts that should never be suspended.

At some point in the future we may choose to integrate the functionality here into role_propagation, making Sherlock responsible for *all* Firecloud account management, instead of just suspension.
