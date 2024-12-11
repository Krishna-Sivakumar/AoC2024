#!/usr/bin/python3
# requires python3.4+ to support Path().touch()

from sys import argv
from pathlib import Path


def main():
    try:
        problem_name = argv[1]
    except:
        print("no problem name provided.")
        return

    if Path(f'{problem_name}.go').exists():
        print("problem file already exists.")
        return

    with open("common/template") as template:
        with open(f"{problem_name}.go", "w") as file:
            file.write(template.read())

        Path(f"inputs/{problem_name}a.in").touch()
        Path(f"inputs/{problem_name}b.in").touch()


if __name__ == "__main__":
    main()
