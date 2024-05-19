import matplotlib.pyplot as plt
import pandas as pd
import argparse

parser = argparse.ArgumentParser()
parser.add_argument("densities", type=str)
parser.add_argument("ratios", type=str)
args = parser.parse_args()

densities = list(map(int, args.densities.split(",")))
ratios = list(map(float, args.ratios.split(",")))

data = {
    "Density": densities,
    "Ratio": ratios,
}

df = pd.DataFrame(data)

df["NormalizedRatio"] = df["Ratio"] / df["Ratio"].max()

df = df.sort_values("Density")

plt.plot(df["Density"], df["NormalizedRatio"])
plt.xlabel("Density")
plt.ylabel("Normalized Ratio")
plt.title("Density vs Normalized Ratio")

max_ratio = df["NormalizedRatio"].max()
max_density = df["Density"][df["NormalizedRatio"].idxmax()]
plt.scatter(max_density, max_ratio, color="red")
plt.annotate(
    f"Max ratio Density {max_density}",
    (max_density, max_ratio),
    textcoords="offset points",
    xytext=(-10, -10),
    ha="center",
)

plt.show()
