import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
from matplotlib.font_manager import FontProperties

# 设置中文字体
font = FontProperties(fname='/System/Library/Fonts/STHeiti Light.ttc', size=10)

# 读取数据
data = []
with open('res.txt', 'r') as file:
    for line in file:
        parts = line.strip().split()
        if len(parts) == 5:
            data.append(parts)

# 创建DataFrame
df = pd.DataFrame(data, columns=['进程编号', '进程ID', '进度', '进度百分比', '时间戳'])
df['进程编号'] = df['进程编号'].astype(int)
df['进度'] = df['进度'].astype(float)
df['时间戳'] = pd.to_datetime(df['时间戳'])

# 计算相对时间
start_time = df['时间戳'].min()
df['相对时间'] = (df['时间戳'] - start_time).dt.total_seconds()

# 设置绘图风格
sns.set_style("whitegrid")
sns.set_palette("husl")

# 图1：横坐标为相对时间，纵坐标为进程编号
plt.figure(figsize=(15, 8))
for process in df['进程编号'].unique():
    process_data = df[df['进程编号'] == process]
    plt.scatter(process_data['相对时间'], process_data['进程编号'], label=f'进程 {process}', s=10)

plt.xlabel('相对时间 (秒)', fontproperties=font)
plt.ylabel('进程编号', fontproperties=font)
plt.title('进程执行时间图', fontproperties=font)
plt.legend(bbox_to_anchor=(1.05, 1), loc='upper left', prop=font)
plt.tight_layout()
plt.savefig('img/进程执行时间图.png', dpi=300, bbox_inches='tight')
plt.close()

# 图2：横坐标为相对时间，纵坐标为进度
plt.figure(figsize=(15, 8))
for process in df['进程编号'].unique():
    process_data = df[df['进程编号'] == process]
    plt.scatter(process_data['相对时间'], process_data['进度'], label=f'进程 {process}', s=10)

plt.xlabel('相对时间 (秒)', fontproperties=font)
plt.ylabel('进度', fontproperties=font)
plt.title('进程进度图', fontproperties=font)
plt.legend(bbox_to_anchor=(1.05, 1), loc='upper left', prop=font)
plt.tight_layout()
plt.savefig('img/进程进度图.png', dpi=300, bbox_inches='tight')
plt.close()

print("图表已生成：进程执行时间图.png 和 进程进度图.png")
