const base = '/'
const routers = [
  {
    name: '添加博客',
    child: [
      {
        name: '添加博客',
        url: '/article/add'
      }
    ]
  }
]
 
export { routers, base }