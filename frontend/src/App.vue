<template>
  <div id="app">
    <el-container class="layout-container">
      <el-aside width="200px" class="sidebar">
        <div class="logo">
          <h2>🪒 Razor Blade</h2>
        </div>
        <el-menu
          :default-active="$route.path"
          class="sidebar-menu"
          router
          :collapse="false"
        >
          <el-menu-item index="/">
            <el-icon><House /></el-icon>
            <span>仪表板</span>
          </el-menu-item>
          <el-menu-item index="/usage-records">
            <el-icon><DocumentCopy /></el-icon>
            <span>使用记录</span>
          </el-menu-item>
          <el-menu-item index="/razors">
            <el-icon><Operation /></el-icon>
            <span>剃须刀管理</span>
          </el-menu-item>
          <el-menu-item index="/blades">
            <el-icon><Coin /></el-icon>
            <span>刀片管理</span>
          </el-menu-item>
          <el-menu-item index="/statistics">
            <el-icon><TrendCharts /></el-icon>
            <span>统计分析</span>
          </el-menu-item>
        </el-menu>
      </el-aside>

      <el-container>
        <el-header class="header">
          <div class="header-content">
            <h3>{{ getPageTitle() }}</h3>
          </div>
        </el-header>

        <el-main class="main-content">
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const pageTitle = computed(() => {
  const titleMap: Record<string, string> = {
    '/': '仪表板',
    '/usage-records': '使用记录',
    '/razors': '剃须刀管理',
    '/blades': '刀片管理',
    '/statistics': '统计分析'
  }
  return titleMap[route.path] || '未知页面'
})

const getPageTitle = () => pageTitle.value
</script>

<style scoped>
.layout-container {
  height: 100vh;
}

.sidebar {
  background-color: #2c3e50;
  color: white;
}

.logo {
  padding: 20px;
  text-align: center;
  border-bottom: 1px solid #34495e;
}

.logo h2 {
  margin: 0;
  color: white;
}

.sidebar-menu {
  border: none;
  background-color: transparent;
}

.sidebar-menu .el-menu-item {
  color: #bdc3c7;
}

.sidebar-menu .el-menu-item:hover {
  background-color: #34495e;
  color: white;
}

.sidebar-menu .el-menu-item.is-active {
  background-color: #3498db;
  color: white;
}

.header {
  background-color: #fff;
  border-bottom: 1px solid #e0e0e0;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.1);
}

.header-content {
  display: flex;
  align-items: center;
  height: 100%;
}

.header-content h3 {
  margin: 0;
  color: #2c3e50;
}

.main-content {
  background-color: #f5f5f5;
  padding: 20px;
}
</style>

<style>
body {
  margin: 0;
  font-family: 'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', '微软雅黑', Arial, sans-serif;
}

#app {
  height: 100vh;
}
</style>