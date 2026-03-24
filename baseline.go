def load_metrics(path):
    return pd.read_csv(path, parse_dates=["timestamp"])
def detect_anomalies(df):
    cpu_thr = 0.8
    mem_thr = 0.85
    lat_thr = 500
    df["cpu_anom"] = df["cpu_usage"] > cpu_thr
    df["mem_anom"] = df["memory_usage"] > mem_thr
    df["lat_anom"] = df["request_latency_ms"] > lat_thr
    df["is_anomaly"] = df[["cpu_anom", "mem_anom", "lat_anom"]].any(axis=1)
    return df
def main():
    metrics_path = Path("metrics_baseline.csv")
    df = load_metrics(metrics_path)
    df = detect_anomalies(df)
    df.to_csv("metrics_with_flags.csv", index=False)
    anomalies = df[df["is_anomaly"]]
    anomalies[[
        "timestamp",
        "cpu_usage",
        "memory_usage",
        "request_latency_ms",
        "is_anomaly"
    ]].to_csv("baseline_anomalies.csv", index=False)
    print(f"Total points: {len(df)}, anomalies: {len(anomalies)}")
if __name__ == "__main__":
    main()
