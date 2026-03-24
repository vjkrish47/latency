# latency
**Hybrid Telemetry Fusion for Early Detection of Systemwide Failures**
* Author: Vijaya Krishna Namala
* Published In : International Journal of Intelligent Systems and Applications in Engineering (IJISAE)
* Publication Date: Jan 2023
* E-ISSN: 2147-6799

**Abstract:**
Modern distributed systems generate diverse telemetry data, but isolated analysis of metrics, logs, and traces delays detection of systemwide failures. This work proposes a hybrid telemetry fusion framework that integrates multi source data into a unified representation for improved observability. By enabling cross layer correlation and temporal alignment, the approach captures early indicators of cascading failures more effectively. The framework enhances early detection capability and supports proactive identification of system instability in large scale distributed environments.

**Key Contributions:**
* **Hybrid Telemetry Fusion Framework Design:**\
Developed a unified framework that integrates metrics, logs, traces, and network signals to enable early detection of systemwide failures in distributed environments.
* **Cross Layer Correlation Mechanism:**\
Designed techniques for telemetry alignment, temporal correlation, and semantic enrichment to capture relationships across multiple system layers.
* **Machine Learning Driven Detection Model:**\
Implemented a fusion based detection approach using machine learning to identify early anomaly patterns and fault propagation signals.
* **Experimental Validation Across Cluster Sizes:**\
Conducted performance evaluation on clusters with 3, 5, 7, 9, and 11 nodes, demonstrating reduced detection latency compared to baseline methods.

**Relevance & Real World Impact**
* **Early Failure Detection Capability :**\
The fusion framework enables earlier identification of systemwide failures by correlating multi source telemetry signals before visible degradation occurs.
* **Reduced Detection Latency :**\
Significant reduction in mean detection latency improves system responsiveness and supports proactive failure mitigation in distributed environments.
**Improved System Reliability :**\
By detecting anomalies at an early stage, the approach helps prevent cascading failures and enhances stability across large scale distributed systems.
**Academic Recognition :**\
Provides a strong foundation for research in observability, anomaly detection, and telemetry fusion within distributed system environments.
**Educational Impact :**\
The work supports academic learning and research by demonstrating the importance of multi source telemetry integration for advanced system monitoring.

**Experimental Results (Summary)**:

  | Nodes | Database DNS Lookup Time - Cached (ms) | Host Table Lookup Time (ms) | Improvement (%) |
  |-------|----------------------------------------| ----------------------------| ----------------|
  | 3     | 0.8                                    | 0.1                         | 87.5            |
  | 5     | 0.8                                    | 0.15                        | 81.3            |
  | 7     | 0.9                                    | 0.2                         | 77.8            |
  | 9     | 0.9                                    | 0.25                        | 72.2            |
  | 11    | 0.9                                    | 0.3                         | 66.7            |

**Citation** \
OPTIMIZING DNS LOOKUP LATENCY FOR ENHANCED PERFORMANCE IN CLUSTERED SYSTEMS. \
Kanagalakshmi Murugan \
International Journal on Science and Technology \
E-ISSN-2229-7677 \
**License** \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijsat.org/ \
**Author Contact** \
**LinkedIn**: http://linkedin.com/in/kanagalakshmi-murugan-12b278199 | **Email**: kanagalakshmi2004@gmail.com
