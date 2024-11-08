import 'dart:io';
import 'dart:io';
import 'package:path/path.dart' as path;

void main() async {
  var mdfDate = '2024-11-06';
  final modifiedAfter = DateTime.parse(mdfDate); // 指定日期
  final sourceDirPath = 'C:\\0prjNtpc\\digital-bet-service'; // 主目录路径
  final targetDirPath = '/bkPrjDBS241108'; // 输出目录路径
   var extnames="java,xml";
  await copyModifiedFiles(modifiedAfter,extnames, sourceDirPath, targetDirPath);
}
/// 复制修改日期在指定日期之后的文件到输出目录，并保留原有的目录结构
Future<void> copyModifiedFiles(
    DateTime modifiedAfter,String extnames, String sourceDirPath, String targetDirPath) async {
  final sourceDir = Directory(sourceDirPath);
  final targetDir = Directory(targetDirPath);

  if (!await sourceDir.exists()) {
    throw Exception("主目录不存在：$sourceDirPath");
  }
  await targetDir.create(recursive: true);

  await for (var entity in sourceDir.list(recursive: true)) {
    if (entity is File) {
      final modifiedDate = await entity.lastModified();
      if (modifiedDate.isAfter(modifiedAfter)) {
        print(entity);
        // 获取扩展名
        final fileExtension = path.extension(entity.path);
        if(fileExtension=="")
          continue;
         var filext=fileExtension.substring(1);
        if(IsInSet(filext,extnames)){
          final relativePath = entity.path.substring(sourceDir.path.length);
          final targetFile = File('${targetDir.path}$relativePath');
          await targetFile.parent.create(recursive: true);
          await entity.copy(targetFile.path);
          print("已复制文件: ${entity.path} 到 ${targetFile.path}");

        }

      }  // end mdfIsafter
    }  //end isFil
  }  //for list
}   //end fun


/**
 * 判断s是否在extnames里面，extnames是个逗号分割字符串
 */
bool IsInSet(String s, String extnames) {
// 将 extnames 按逗号分割成一个列表，并去掉每个项的空格
  List<String> extList = extnames.split(',').map((e) => e.trim()).toList();

  // 判断 s 是否在分割后的列表中
  return extList.contains(s);


}

 
