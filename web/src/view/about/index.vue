<template>
  <div class="mt-2">
    <el-row :gutter="10">
      <el-col :span="22" :offset="1">
        <el-card>
          <div>
            <el-row>
              <el-col :span="4">
                <a href="https://github.com/flipped-aurora/gin-vue-admin">
                  <img
                    class="org-img dom-center"
                    src="@/assets/logo.png"
                    alt="gin-vue-admin"
                  >
                </a>
              </el-col>
              <el-col :span="18" >
                <p style="font-size: larger;">
                  在当今的数字化时代，网络安全已成为防御体系中的关键一环。随着网络攻击手段日益狡猾，传统安全防御机制已不再能够满足当前的需求。我们需要诱导网络攻击并尝试反制溯源，应对这一挑战，我们开发了创新性的“多维协议蜜罐系统”。该系统不仅整合了多种网络协议蜜罐、数据库蜜罐和工控蜜罐，还引入了虚拟主机功能，通过Docker技术实现虚拟化，能在单一主机上部署众多蜜罐，甚至模拟出大量虚拟网络设备。
                </p>
              <br/>
                <p style="font-size: larger;"
                  >本系统采用了前后端分离的技术架构，其后台管理系统界面直观易用，允许安全管理员轻松监控和管理各类蜜罐的运行状况，新建和删除蜜罐。通过实时捕捉并分析攻击者行为，我们的系统能够提供深入的攻击见解，帮助安全团队更全面地理解潜在安全威胁，并优化其防御策略。 </p
                ><br/><p style="font-size: larger;"
                  >我们的蜜罐系统不仅仅是一个被动防御工具，它还能主动搜集攻击者信息。凭借精确的日志记录和智能分析功能，它为用户提供了全面的安全保障，并为未来安全研究积累了宝贵的数据。通过模拟真实网络环境，我们的系统能够诱导攻击者暴露其攻击手法，甚至帮助追踪到攻击者本身，为全球网络安全作出贡献。</p
                >
              </el-col>
            </el-row>

          </div>
        </el-card>
      </el-col>

    </el-row>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Commits, Members } from '@/api/github'
import { formatTimeToStr } from '@/utils/date'
const page = ref(0)

defineOptions({
  name: 'About'
})

const loadMore = () => {
  page.value++
  loadCommits()
}

const dataTimeline = ref([])
const loadCommits = () => {
  Commits(page.value).then(({ data }) => {
    data.forEach((element) => {
      if (element.commit.message) {
        dataTimeline.value.push({
          from: formatTimeToStr(element.commit.author.date, 'yyyy-MM-dd'),
          title: element.commit.author.name,
          showDayAndMonth: true,
          message: element.commit.message,
        })
      }
    })
  })
}

const members = ref([])
const loadMembers = () => {
  Members().then(({ data }) => {
    members.value = data
    members.value.sort()
  })
}

loadCommits()
loadMembers()

</script>

<style scoped>
.load-more {
  margin-left: 120px;
}

.avatar-img {
  float: left;
  height: 40px;
  width: 40px;
  border-radius: 50%;
  -webkit-border-radius: 50%;
  -moz-border-radius: 50%;
  margin-top: 15px;
}

.org-img {
  height: 150px;
  width: 150px;
}


.dom-center {
  margin-left: 50%;
  transform: translateX(-50%);
}
</style>
