### 关联查询
[参考](https://github.com/jinzhu/gorm/issues/2193)

```go
我现在是查完list再range去获取related
`
func (p *Post) GetUserPostsByUserId(userId uint64, page int, size int) *[]Post {
posts := []Post{}
if err := DB().Where("user_id = ?", userId).Offset((page - 1) * size).Limit(size).Find(&posts).Error; err != nil {
log.Debugf("Get user posts error: %v", err)
return nil
}
for key, post := range posts {
if err := DB().Model(&post).Related(&post.User).Error; err != nil {
log.Debugf("Post user related error: %v", err)
}
posts[key] = post
}

return &posts
}
```