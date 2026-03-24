import pandas as pd
from pathlib import Path
from sklearn.ensemble import RandomForestClassifier
from sklearn.model_selection import train_test_split
from sklearn.metrics import classification_report

def load_telemetry(metrics_path, logs_path):
    m = pd.read_csv(metrics_path, parse_dates=["timestamp"])
    l = pd.read_csv(logs_path, parse_dates=["timestamp"])
    l_agg = l.groupby("timestamp")["log_severity"].max().reset_index()
    return m.merge(l_agg, on="timestamp", how="left").fillna({"log_severity": 0})

def train_model(df):
    features = ["cpu_usage", "memory_usage", "request_latency_ms", "log_severity"]
    X = df[features]
    y = df["failure_label"]
    X_train, X_test, y_train, y_test = train_test_split(
        X, y, test_size=0.2, shuffle=False, random_state=42
    )
    clf = RandomForestClassifier(n_estimators=100, random_state=42)
    clf.fit(X_train, y_train)
    y_pred = clf.predict(X_test)
    print(classification_report(y_test, y_pred, digits=3))
    df["predicted_failure"] = clf.predict(X)
    return clf, df

def main():
    base = Path(".")
    df = load_telemetry(base / "metrics_fusion.csv", base / "logs_fusion.csv")
    model, df_scored = train_model(df)
    df_scored.to_csv("fusion_predictions.csv", index=False)
    print(f"Scored rows: {len(df_scored)}")

if __name__ == "__main__":
    main()
