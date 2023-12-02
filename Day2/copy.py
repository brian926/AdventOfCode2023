import math
import pathlib
import sys

from collections import defaultdict
from typing import TextIO

sys.path.append(str(pathlib.Path(__file__).resolve().parents[3] / 'lib' / 'python'))

#import aoc

Pull = tuple[int, int, int]
Games = dict[int, list[Pull]]

def parse_input(f: TextIO) -> Games:
  games = {}

  for line in f.readlines():
    (game, results) = line.strip().split(": ")

    game_id = int(game[5:])
    games[game_id] = []

    for result in results.split("; "):
      pull = defaultdict(int, {
        cube: int(count)
        for count, cube in (cubes.split(" ") for cubes in result.split(", "))
      })
      games[game_id].append((pull["red"], pull["green"], pull["blue"]))

  return games

def power(pull: Pull) -> int:
  return math.prod(x for x in pull)

def possible_games(games: Games, cubes: Pull) -> list[int]:
  return [
    id
    for id, pulls in games.items()
    if all(
      all(max_cube >= cube for (max_cube, cube) in zip(cubes, pull))
      for pull in pulls
    )
  ]

def min_pulls(games: Games) -> list[Pull]:
  return [tuple(max(x) for x in zip(*pulls)) for pulls in games.values()]  # type: ignore

def run() -> None:
  with open(aoc.inputfile('input.txt')) as f:
     games = parse_input(f)

  game_ids = possible_games(games, (12, 13, 14))
  powers = [power(pull) for pull in min_pulls(games)]

  print(f"Sum of possible game IDs: {sum(game_ids)}")
  print(f"Sum of minimum powers: {sum(powers)}")

if __name__ == '__main__':
  run()
  sys.exit(0)