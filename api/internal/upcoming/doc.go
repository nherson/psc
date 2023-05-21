/*
This is a monolithic package that includes a db client, tapology client, fightinsider client,
and fuzzy matcher. It synthesizes all data to entirely construct a tree of upcoming fight
data and odds, performing a best-effort association of existing fighters to their upcoming
fights and odds.

The routines in this package will, in a single transaction:
- purge all upcoming events, fights, and temporary fighter nodes
- attempt to perform all needed associations between DB data and tapology+fightinsider data
- re-create all upcoming events, fights and temporary fighters nodes, driven by updated data from clients

Methods provided to do the above both with prompts for CLI usage (picklist when no clear match), and automated for API usage
*/
package upcoming
