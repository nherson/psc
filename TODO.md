## Base work

- ~~Ent models~~
- ~~Conversion from ufc struct to Ent model~~
- ~~Bulk upsert of historical data (mage target)~~
- ~~API: upcoming~~ REQUIRES TESTING
- ~~Mage: historical import (import:all and import:event XXXX)~~
- ~~Add timestamps to events for when they happened~~
- ~~Clean up multiple connection functions (and rollback function)~~
- ~~Setup Chakra UI basics~~
- ~~Build UI parity with old site~~
- ~~Refactor fightodds prompt+fuzzy-matcher into its own package, to be reused by tapology~~
- ~~Cron to remove upcoming fights and events that are no longer "upcoming"~~
- UI for upcoming event and fight info
- ~~Personal data mangling endpoint, hands back a CSV~~
  - ~~How I like the data~~
- ~~Public data endpoint~~
  - ~~CSV~~
  - ~~For every upcoming fighter, one row for each past fight~~
    - ~~Name, sig strikes, takedowns, knockdowns, control time, odds, score, date~~
  - ~~Add date of previous fight to CSV data~~
- Logging
- Use React env vars to determine URL for both twirp and stat clients
- Add "last updated" info somewhere near the stat downlaod button
- "Fighter Alias" table to stash other names for fighters to bypass future similarity-based matches
- Fix layout on mobile
- ~~HTTPS redirect~~

## Monthly process

- ~~Load in data from previous month (TODO: mage target)~~
- ~~Handle upcoming events and fights~~
  <!-- * Create new fighter nodes (TODO: make mage target for this) -->
  - ~~Register upcoming fights and events, make new fighters as needed, wire in odds from odds site~~
    - ~~Would be cool if this had some sort of "preview" output before committing to the DB~~

## Nice to haves

- API to update upcoming events and fights
  - Make it an API endpoint and hit it on a cron?
- ~~Add "corner" to FightResult for a fighter, for visual consistency (Red left, Blue right)~~
- Include venue information for events
- Use Grid to equally space the fighter result cards

## Upcoming Import Process (implemented!)

- For upcoming DB loading
  - Wipe all Upcoming\* records and re-create in a single transaction
  - Not possible to "upsert" since upcoming fights have no unique identifier and the card is constantly changing
  - To associate fighters, go through these cases in order
    - If tapology_id set on fighter and matches, done!
      - Can tapology IDs change?
        - If yes, maybe need to store just the number part of the id (e.g. `123456-jon-doe-the-killa` vs `123456`)
          - DONT DO THIS: found at least one fighter with no number prefix, just their name as the id
    - If fuzzy match is above certain threshold, associate tapology_id
    - Else
      - If automated process, create temporary fighter record (need bool on fighter schema)
      - If mage task, prompt fuzzy match best options
        - If none, create temporary record
