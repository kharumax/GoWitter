# GoWitter

## やること
①会員登録機能  
　ーDBテーブル作成  
　ーフォーム情報を取得  
　ーDB格納
②ログイン機能
　ーフォーム情報を取得  
　ーDB確認＆セッション情報作成  
　ーCookieに格納＆リダイレクト  
　ーログイン保持作成  
　ーログアウト  
③投稿機能  
　ー投稿一覧＆詳細  
　ー新規投稿  
　　・ログイン有無の確認  
　　・CSRF対策  
　ー投稿編集＆削除  
　　・ログインユーザー認証  
　　・CSRF対策  
④いいね機能  
  ーいいね済みか確認
  ーログイン
  ーいいね機能（Ajax)
  ーいいね取り消し（Ajax)
⑤コメント機能
　ーログイン
　ーコメント機能
⑥フォロー機能
　ーフォロー済み確認＆ログイン
　ーフォロー

## DB
-User
 -id
 -email
 -name
 -profile_image
 -description
-Post
 -id
 -content
 -image
 -user_id
-Like
 -id
 -user_id
 -post_id
-Comment
 -id
 -user_id
 -post_id
 -text

## 機能
1. SignUp  
2. Login/Logout  
3. Post  
   -new   
   -list  
   -show  
   -edit  
   -delete  
4. Like  
   -new  
   -delete  
5. Comment  
   -new  
   -delete  
6. Follow  
   -new  
   -delete  
   
   