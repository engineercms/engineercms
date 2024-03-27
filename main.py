# 这是一个示例 Python 脚本。

# 按 Shift+F10 执行或将其替换为您的代码。
# 按 双击 Shift 在所有地方搜索类、文件、工具窗口、操作和设置。

# FREECADPATH = '/usr/lib/freecad-python3/lib/'  # path to your FreeCAD.so or FreeCAD.pyd file,
# for Windows you must either use \\ or / in the path, using a single \ is problematic
# FREECADPATH = 'C:\\FreeCAD\\bin'
FREECADPATH = 'D:/Program Files/FreeCAD 0.21/bin'
# import Blender, sys
import sys
import argparse

sys.path.append(FREECADPATH)


# import FreeCAD
def import_fcstd(inputfile, parameter, outputfile):
    try:
        import FreeCAD
    except ValueError:
        print(
            'Error%t|FreeCAD library not found. Please check the FREECADPATH variable in the import script is correct')
    # Blender.Draw.PupMenu(
    #     'Error%t|FreeCAD library not found. Please check the FREECADPATH variable in the import script is correct')
    else:
        import Part
        doc = FreeCAD.open(inputfile)
        ss = doc.getObject('Spreadsheet')
        ss.set('C5', parameter)
        FreeCAD.ActiveDocument.recompute()

        __objs__ = []
        objects = doc.Objects

        for ob in objects:
            # print(ob.Name+','+ob.TypeId)
            if ob.Name[:4] == 'Body' or ob.Name[:4] == 'Part':
                __objs__.append(ob)

        if hasattr(Part, "exportOptions"):
            options = Part.exportOptions(u'D:/' + outputfile)
            Part.export(__objs__, u'D:/' + outputfile, options)
        else:
            Part.export(__objs__, u'D:/' + outputfile)

    del __objs__


# >>> __objs__ = []
# >>> __objs__.append(FreeCAD.getDocument("_33").getObject("Pad001"))
# >>> import Part
# >>> if hasattr(Part, "exportOptions"):
# >>>     options = Part.exportOptions(u"D:/11.stp")
# >>>     Part.export(__objs__, u"D:/11.stp", options)
# >>> else:
# >>>     Part.export(__objs__, u"D:/11.stp")
# >>>
# >>> del __objs__


# def main():
#     Blender.Window.FileSelector(import_fcstd, 'IMPORT FCSTD',
#                                 Blender.sys.makename(ext='.fcstd'))

# This lets you import the script without running it
if __name__ == '__main__':
    # parser.add_argument('-u', '--user', dest='User', type=str,default='root', help='target User')
    # parser.add_argument('-s', '--sex', dest='Sex', type=str, choices=['男', '女'], default='男', help='target Sex')
    # parser.add_argument('-n', '--number', dest='Num', nargs=2, required=True,type=int, help='target Two Numbers')

    parser = argparse.ArgumentParser(description="读取FreeCAD文件(*.FCStd)根据输入参数修改模型导出stp格式模型文件")
    parser.add_argument('-i', '--input_file', dest='input_file', required=True, type=str,
                        help='target InputFile FreeCAD FCStd')
    parser.add_argument('-p', '--parameter_file', dest='parameter_file', required=True, type=str,
                        help='target ParameterFile text')
    parser.add_argument('-o', '--output_file', dest='output_file', required=True, type=str,
                        help='target OutputFile stp')

    # ！！！！坑2：-m 是绝对不行的，不能用这个字母，奇怪！！
    # dir_path = os.path.dirname(os.path.realpath(__file__))   #目录路径
    # input_file=os.path.join(dir_path, 'damrst.rst')
    # mid_file=os.path.join(dir_path, 'damrst.vtk')
    # output_file=os.path.join(dir_path, 'damrst.vtp')
    # sys.exit(read_ansys(input_file,mid_file,output_file))

    args = parser.parse_args()
    sys.exit(import_fcstd(args.input_file, args.parameter_file, args.output_file))

# def print_hi(name):
#     # 在下面的代码行中使用断点来调试脚本。
#     print(f'Hi, {name}')  # 按 Ctrl+F8 切换断点。
#
#
# # 按装订区域中的绿色按钮以运行脚本。
# if __name__ == '__main__':
#     print_hi('PyCharm')

# 访问 https://www.jetbrains.com/help/pycharm/ 获取 PyCharm 帮助
