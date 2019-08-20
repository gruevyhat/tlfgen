Character Generator for the Laundry Files RPG
==============================================

[![Go Report Card](https://goreportcard.com/badge/github.com/gruevyhat/tlfgen)](https://goreportcard.com/report/github.com/gruevyhat/tlfgen)

Character generator for the Laundry Files RPG.

## Installation

    $ git clone https://github.com/gruevyhat/tlfgen
    $ cd tlfgen
    $ make

## Usage

    The Laundry Files Character Generator

    Usage: tlfgen [options]

    Options:
      -n, --name=<str>          The character's full name.
      -g, --gender=<str>        The character's gender.
      -a, --age=<int>           The character's age.
      -p, --personality=<str>   The character's personality type.
      -A, --assignment=<str>    The character's assignment.
      -P, --profession=<str>    The character's profession.
      -S, --skill-points=<str>  The character's bonus skill points.
      -b, --attr-bonus=<str>    The character's attribute bonus ("smart", "tough", or "mystical").
      -s, --seed=<hex>          Character generation signature.
      --log-level=<str>         One of {INFO, WARNING, ERROR}. [default: ERROR]
      -h --help
      --version

## Example

    $ tflgen -A "Computational Demonology" -P "Parapsychologist" -p "Thinker" -S 200 | jq -c '.'
    {
        "name": "Jacob Garcia",
        "age": 19,
        "gender": "Male",
        "personality_type": "Thinker",
        "profession": "Parapsychologist",
        "assignment": "Computational Demonology",
        "wealth": "Average",
        "base": {
            "charisma": 18,
            "constitution": 15,
            "dexterity": 10,
            "education": 12,
            "intelligence": 15,
            "power": 11,
            "size": 15,
            "strength": 8
        },
        "derived": {
            "agility": 50,
            "damage_bonus": "None",
            "damage_mod": 0,
            "effort": 40,
            "endurance": 75,
            "experience_bonus": 7,
            "hit_points": 15,
            "idea": 75,
            "know": 60,
            "luck": 55,
            "major_wound_level": 7,
            "move": 10,
            "sanity": 55
        },
        "skills": {
            "Appraise": 44,
            "Athletics": 10,
            "Bargain": 28,
            "Brawl": 27,
            "Bureaucracy": 12,
            "Climb": 40,
            "Command": 5,
            "Computer Use: Hacking": 18,
            "Computer Use: Magic": 23,
            "Computer Use: Programming": 11,
            "Demolition": 3,
            "Disguise": 34,
            "Dodge": 23,
            "Drive: Automobile": 22,
            "Etiquette": 6,
            "Fast Talk": 32,
            "Fine Manipulation": 10,
            "Firearm: Exotic": 32,
            "Firearm: Pistol": 27,
            "First Aid": 37,
            "Gaming": 12,
            "Grapple": 30,
            "Hide": 38,
            "Insight": 50,
            "Jump": 27,
            "Knowledge: Accounting": 17,
            "Knowledge: Espionage": 9,
            "Knowledge: Law": 12,
            "Knowledge: Occult": 64,
            "Knowledge: Politics": 11,
            "Knowledge: Streetwise": 37,
            "Language: Own": 79,
            "Listen": 86,
            "Navigate": 11,
            "Perform": 7,
            "Persuade": 19,
            "Research": 77,
            "Ride": 5,
            "Science: Mathematics": 21,
            "Science: Thaumaturgy": 14,
            "Sense": 58,
            "Sleight of Hand": 9,
            "Sorcery": 13,
            "Spot": 83,
            "Status": 17,
            "Stealth": 66,
            "Strategy": 5,
            "Swim": 28,
            "Teach": 12,
            "Technology Use: Other": 29,
            "Throw": 28,
            "Track": 13
        },
        "seed": "15bc33ec5a96a2bc"
    }

