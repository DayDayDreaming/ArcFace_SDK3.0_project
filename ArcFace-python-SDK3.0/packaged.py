import cv2
from arcface.engine import *
import os
import sys


#激活接口,首次需联网激活
def Initialization():
    APPID = b'GkhGbHoWtKqfEZqswAuo5QziHsxoDg2pKD9A4Vr3mttz'
    SDKKey = b'9AyFe6dvTcHiEXzs8JujKcCvgo5dxYJ1RGMyh61jzmYg'

    # 激活接口,首次需联网激活
    res = ASFOnlineActivation(APPID, SDKKey)
    if (MOK != res and MERR_ASF_ALREADY_ACTIVATED != res):
        print("ASFActivation fail: {}".format(res))
    else:
        print("ASFActivation sucess: {}".format(res))



#与人脸库"./asserts"文件夹中的照片进行比对,检测是否注册过，能否通过人脸验证
def FaceComparation():
    Initialization()
    # 获取人脸识别引擎
    face_engine = ArcFace()

    # 需要引擎开启的功能
    mask = ASF_FACE_DETECT | ASF_FACERECOGNITION | ASF_AGE | ASF_GENDER | ASF_FACE3DANGLE | ASF_LIVENESS | ASF_IR_LIVENESS

    # 初始化接口
    res = face_engine.ASFInitEngine(ASF_DETECT_MODE_IMAGE, ASF_OP_0_ONLY, 30, 10, mask)
    if (res != MOK):
        print("ASFInitEngine fail: {}".format(res))
    else:
        print("ASFInitEngine sucess: {}".format(res))


    # 检测摄像头中的人脸
    capture = cv2.VideoCapture(0)
    while (True):
        ref, frame = capture.read()
        #cv2.imshow("1", frame)
        c = cv2.waitKey(1) & 0xff

        res, detectedFaces1 = face_engine.ASFDetectFaces(frame)

        if detectedFaces1.faceNum != 1:
            continue

        # 设置活体置信度 SDK内部默认值为 IR：0.7 RGB：0.75（无特殊需要，可以不设置）
        threshold = ASF_LivenessThreshold()
        threshold.thresholdmodel_BGR = 0.75
        threshold.thresholdmodel_IR = 0.7

        face_engine.ASFSetLivenessParam(threshold)

        # RGB图像属性检测 注意:processMask中的内容必须在初始化引擎 时指定的功能内
        processMask = ASF_AGE | ASF_GENDER | ASF_FACE3DANGLE | ASF_LIVENESS

        res = face_engine.ASFProcess(frame, detectedFaces1, processMask)

        if res == MOK:
            # 获取年龄
            res, ageInfo = face_engine.ASFGetAge()
            if (res != MOK):
                print("ASFGetAge fail: {}".format(res))
            else:
                print("Age: {}".format(ageInfo.ageArray[0]))

            # 获取性别
            res, genderInfo = face_engine.ASFGetGender()
            if (res != MOK):
                print("ASFGetGender fail: {}".format(res))
            else:
                print("Gender: {}".format(genderInfo.genderArray[0]))

            # 获取3D角度
            res, angleInfo = face_engine.ASFGetFace3DAngle()
            if (res != MOK):
                print("ASFGetFace3DAngle fail: {}".format(res))
            else:
                print("3DAngle roll: {} yaw: {} pitch: {}".format(angleInfo.roll[0],
                                                                  angleInfo.yaw[0], angleInfo.pitch[0]))

            # 获取RGB活体信息
            res, rgbLivenessInfo = face_engine.ASFGetLivenessScore()
            if (res != MOK):
                print("ASFGetLivenessScore fail: {}".format(res))
            else:
                print("RGB Liveness: {}".format(rgbLivenessInfo.isLive[0]))
        else:
            print("ASFProcess fail: {}".format(res))
            if (res == 81925) :
                #大概率是识别角度问题
                print("请重新识别")
                capture.release()
                face_engine.ASFUninitEngine()
                return


        #isliveness为1时表明通过了rbg活体检测，跳出循环进行面部特征提取
        if (rgbLivenessInfo.isLive[0] == 1):
            break

    capture.release()

    print(detectedFaces1)
    if res == MOK:
        single_detected_face1 = ASF_SingleFaceInfo()
        single_detected_face1.faceRect = detectedFaces1.faceRect[0]
        single_detected_face1.faceOrient = detectedFaces1.faceOrient[0]
        res, face_feature1 = face_engine.ASFFaceFeatureExtract(frame, single_detected_face1)
        if (res != MOK):
            print("ASFFaceFeatureExtract 1 fail: {}".format(res))
    else:
        print("ASFDetectFaces 1 fail: {}".format(res))


    #对"./asserts"文件夹中已经注册的照片进行遍历对比
    for file in os.listdir(r"C:/Users/Assassin/Desktop/ArcFace-python-master/asserts"):
        filepath = r"C:/Users/Assassin/Desktop/ArcFace-python-master/asserts/" + file
        img2 = cv2.imread(filepath)

        res, detectedFaces2 = face_engine.ASFDetectFaces(img2)
        if res == MOK:
            single_detected_face2 = ASF_SingleFaceInfo()
            single_detected_face2.faceRect = detectedFaces2.faceRect[0]
            single_detected_face2.faceOrient = detectedFaces2.faceOrient[0]
            res, face_feature2 = face_engine.ASFFaceFeatureExtract(img2, single_detected_face2)
            if (res == MOK):
                pass
            else:
                print("ASFFaceFeatureExtract 2 fail: {}".format(res))
        else:
            print("ASFDetectFaces 2 fail: {}".format(res))

        # 比较两个人脸的相似度
        res, score = face_engine.ASFFaceFeatureCompare(face_feature1, face_feature2)
        print("相似度:", score)

        #注册库中找到对应的照片
        if (score >= 0.8) :
            print("Verified")
            # 销毁引擎
            face_engine.ASFUninitEngine()
            return 0


    print("Not register")
    face_engine.ASFUninitEngine()




#注册人脸
def Registeration():
    path = sys.argv[1]
    #path = r"C:\Users\Assassin\Desktop\localhost\static\image\guest\3306\1586503509.jpg"
    print(path)

    Initialization()
    # 获取人脸识别引擎
    face_engine = ArcFace()

    # 需要引擎开启的功能
    mask = ASF_FACE_DETECT | ASF_FACERECOGNITION | ASF_AGE | ASF_GENDER | ASF_FACE3DANGLE | ASF_LIVENESS | ASF_IR_LIVENESS

    # 初始化接口
    res = face_engine.ASFInitEngine(ASF_DETECT_MODE_IMAGE, ASF_OP_0_ONLY, 30, 10, mask)
    if (res != MOK):
        print("ASFInitEngine fail: {}".format(res))
    else:
        print("ASFInitEngine sucess: {}".format(res))

    # 检测摄像头中的人脸
    capture = cv2.VideoCapture(0)
    while (True):
        ref, frame = capture.read()
        #cv2.imshow("1", frame)
        c = cv2.waitKey(1) & 0xff

        res, detectedFaces1 = face_engine.ASFDetectFaces(frame)

        if detectedFaces1.faceNum != 1:
            continue
            
        # 设置活体置信度 SDK内部默认值为 IR：0.7 RGB：0.75（无特殊需要，可以不设置）
        threshold = ASF_LivenessThreshold()
        threshold.thresholdmodel_BGR = 0.75
        threshold.thresholdmodel_IR = 0.7

        face_engine.ASFSetLivenessParam(threshold)

        # RGB图像属性检测 注意:processMask中的内容必须在初始化引擎 时指定的功能内
        processMask = ASF_AGE | ASF_GENDER | ASF_FACE3DANGLE | ASF_LIVENESS

        res = face_engine.ASFProcess(frame, detectedFaces1, processMask)

        if res == MOK:
            # 获取年龄
            res, ageInfo = face_engine.ASFGetAge()
            if (res != MOK):
                print("ASFGetAge fail: {}".format(res))
            else:
                print("Age: {}".format(ageInfo.ageArray[0]))

            # 获取性别
            res, genderInfo = face_engine.ASFGetGender()
            if (res != MOK):
                print("ASFGetGender fail: {}".format(res))
            else:
                print("Gender: {}".format(genderInfo.genderArray[0]))

            # 获取3D角度
            res, angleInfo = face_engine.ASFGetFace3DAngle()
            if (res != MOK):
                print("ASFGetFace3DAngle fail: {}".format(res))
            else:
                print("3DAngle roll: {} yaw: {} pitch: {}".format(angleInfo.roll[0],
                                                                  angleInfo.yaw[0], angleInfo.pitch[0]))

            # 获取RGB活体信息
            res, rgbLivenessInfo = face_engine.ASFGetLivenessScore()
            if (res != MOK):
                print("ASFGetLivenessScore fail: {}".format(res))
            else:
                print("RGB Liveness: {}".format(rgbLivenessInfo.isLive[0]))
        else:
            print("ASFProcess fail: {}".format(res))
            if (res == 81925):
                # 大概率是识别角度问题
                print("请重新识别")
                capture.release()
                face_engine.ASFUninitEngine()
                return 0

        # isliveness为1时表明通过了rbg活体检测，跳出循环进行检验是否注册过
        if (rgbLivenessInfo.isLive[0] == 1):
            break

    capture.release()

    print(detectedFaces1)
    if res == MOK:
        single_detected_face1 = ASF_SingleFaceInfo()
        single_detected_face1.faceRect = detectedFaces1.faceRect[0]
        single_detected_face1.faceOrient = detectedFaces1.faceOrient[0]
        res, face_feature1 = face_engine.ASFFaceFeatureExtract(frame, single_detected_face1)
        if (res != MOK):
            print("ASFFaceFeatureExtract 1 fail: {}".format(res))
    else:
        print("ASFDetectFaces 1 fail: {}".format(res))


    # 对"./asserts"文件夹中已经注册的照片进行遍历对比
    '''for file in os.listdir(r"C:/Users/Assassin/Desktop/ArcFace-python-master/asserts"):
        filepath = r"C:/Users/Assassin/Desktop/ArcFace-python-master/asserts/" + file
        img2 = cv2.imread(filepath)

        res, detectedFaces2 = face_engine.ASFDetectFaces(img2)
        if res == MOK:
            single_detected_face2 = ASF_SingleFaceInfo()
            single_detected_face2.faceRect = detectedFaces2.faceRect[0]
            single_detected_face2.faceOrient = detectedFaces2.faceOrient[0]
            res, face_feature2 = face_engine.ASFFaceFeatureExtract(img2, single_detected_face2)
            if (res == MOK):
                pass
            else:
                print("ASFFaceFeatureExtract 2 fail: {}".format(res))
        else:
            print("ASFDetectFaces 2 fail: {}".format(res))

        # 比较两个人脸的相似度
        res, score = face_engine.ASFFaceFeatureCompare(face_feature1, face_feature2)
        print("相似度:", score)

        if (score >= 0.8) :
            print("has registered")
            # 销毁引擎
            face_engine.ASFUninitEngine()
            return 0

    #命名新注册的人脸
    name = len(os.listdir(r"C:/Users/Assassin/Desktop/ArcFace-python-master/asserts"))+ 1
    name = r"C:/Users/Assassin/Desktop/ArcFace-python-master/asserts//" + str(name) + ".jpg" '''
    #path = r"C:/Users/Assassin/Desktop/localhost/static/image/guest/3306/1586503509.jpg"
    cv2.imwrite(path, frame)
    print("注册完毕")
    # 销毁引擎
    face_engine.ASFUninitEngine()









