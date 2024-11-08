import os
from os import mkdir

# pip install playsound==1.2.2
# skuli not take effk
# pip install
#  pip install Pillow
# //for mod cv2
#//  pip install opencv-python

# pip install --upgrade pip setuptools wheel
# pip install playsound

# pip install screeninfo
# 播放音乐文件
#playsound(r"C:\db\海鸣威-老人与海.mp3")
#playsound("C:\\db\\laor.mp3")
# 检测匹配度
threshold = 0.5
iconMeeting = "C:/db/ggl3.jpg"
import time
from time import sleep
import datetime
from PIL import ImageGrab
import cv2
import numpy as np




from playsound import playsound




from PIL import Image
import mss


def mkdir2024(pngFilepath):
    # 使用 os.makedirs 创建目录，exist_ok=True 表示如果目录已存在，则不会抛出异常
    # 提取文件路径中的目录部分
    directory = os.path.dirname(pngFilepath)
    # 创建目录，包括所有缺失的级联目录
    os.makedirs(directory, exist_ok=True)


def getSecScr():
    # 获取显示器信息
    with mss.mss() as sct:
        # 获取所有屏幕的信息
        monitors = sct.monitors
        if len(monitors) < 2:
            raise Exception("没有找到第二个显示器。")

        # 打印所有显示器的信息
        for i, monitor in enumerate(monitors):
            print(f"Monitor {i}: {monitor['width']}x{monitor['height']} at {monitor['left']},{monitor['top']}")

        # 假设第二个屏幕是第二个显示器（索引为1）
        second_monitor = monitors[2]  # mss 中的 monitors 从 1 开始计数
        bbox = (second_monitor['left'], second_monitor['top'],
                second_monitor['left'] + second_monitor['width'],
                second_monitor['top'] + second_monitor['height'])

        # 截取第二个屏幕的内容
        img = sct.grab(bbox)

        # 转换为 PIL 图像并保存
        img_pil = Image.frombytes("RGB", img.size, img.bgra, "raw", "BGRX")
        pngFilepath = "/zscrpic/second_screen"+generate_filenameTimepart()+".png"
        mkdir2024(pngFilepath)
        img_pil.save(pngFilepath)  # 保存截屏



    print("第二个屏幕的截图已保存为  "+pngFilepath)
    return  pngFilepath

def foreachx():
    # 截屏
    screenshotpath=getSecScr()
    print(screenshotpath)
    #screenshot = ImageGrab.grab()
    # screenshot.save("screenshot.png")

    # 查找图标
    screenshot = cv2.imread(screenshotpath)

    icon = cv2.imread(iconMeeting)
    # ggl2 just with txt ,,ggl onlyh icon
    result = cv2.matchTemplate(screenshot, icon, cv2.TM_CCOEFF_NORMED)
    #print("matchTemplate#result:"+result)
   # print(f"matchTemplate#result: {result.tolist()}")  # 转换为列表再打印
    #print("matchTemplate#result: " + str(result))

    locations = np.where(result >= threshold)

    if locations[0].size > 0:
        print("图标找到，位置：", locations)
        #from playsound import playsound

        # 播放音乐文件
        playsound("C:\\db\\laor.mp3")  # 替换为你的音乐文件路径
    else:
        print("图标未找到。")

def generate_filenameTimepart():
    # 获取当前日期和时间
    now = datetime.datetime.now()
    # 根据日期时间生成文件名，格式为 'YYYY-MM-DD_HH-MM-SS'
    filename = now.strftime("%Y-%m-%d_%H-%M-%S") + ""
    return filename
while True:
    print("这是一个无限循环，每 5 秒打印一次。")
   # time.sleep(5)  # 暂停 5 秒

    foreachx()
    # 等待 5 秒
    sleep(5)


