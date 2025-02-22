<template>
  <el-container style="height: 100vh;">
    <!-- 主要内容区域 -->
    <el-main>
      <Functions />
      <el-row :gutter="20">
        <el-col :span="20">
          <h2>添加商品</h2>
          <el-form ref="addProductFormRef" :model="product" label-width="120px">
            <el-form-item label="商品名称">
              <el-input v-model="product.name" placeholder="请输入商品名称" />
            </el-form-item>
            <el-form-item label="商品类别">
              <el-select v-model="product.categories" multiple placeholder="请选择商品类别" style="width: 100%">
                <el-option v-for="category in categories.slice(1)" :key="category.id" :label="category.name"
                  :value="category.id" />
              </el-select>
            </el-form-item>
            <el-form-item label="商品主图">
              <el-upload class="avatar-uploader" action="/api/uploadImage" :show-file-list="false"
                :on-success="handleMainUpload" :before-upload="beforeAvatarUpload">
                <img v-if="product.img_url" :src="product.img_url" class="avatar">
                <!-- TODO：UI 没生效 -->
                <i v-else class="el-icon-plus avatar-uploader-icon">+</i>
                <template #tip>
                  <div class="upload-tip">建议尺寸：800x800px，支持JPG/PNG格式</div>
                </template>
              </el-upload>
            </el-form-item>


            <el-form-item label="轮播图内容">
              <div class="combination-list">

                <div v-for="(item, index) in product.slider_items" :key="index" class="combination-item">
                  <div class="combination-content">
                    <div class="media-group">
                      <!-- 图片区域 -->
                      <div class="media-container">
                        <template v-if="item.image">
                          <img :src="item.image" class="thumbnail-image">
                          <!-- <div class="remove-btn" @click.stop="removeMedia(index, 'image')">×</div> -->
                        </template>
                        <el-upload v-else action="/api/uploadImage" :show-file-list="false"
                          :on-success="(res) => handleImageUpload(res, index)" :before-upload="beforeImageUpload">
                          <el-button class="upload-btn">+ 添加图片</el-button>
                        </el-upload>
                      </div>

                      <!-- 视频区域 -->
                      <div class="media-container">
                        <template v-if="item.video">
                          <video :src="item.video" class="thumbnail-video"></video>
                          <div class="remove-btn" @click.stop="removeMedia(index, 'video')">删除视频</div>
                        </template>
                        <el-upload v-else-if="item.image" action="/api/uploadVideo" :show-file-list="false"
                          :on-success="(res) => handleVideoUpload(res, index)" :before-upload="beforeVideoUpload">
                          <el-button class="upload-btn">+ 添加视频</el-button>
                        </el-upload>
                        <div v-else class="empty-tip">请先上传图片</div>
                      </div>
                    </div>

                    <div class="remove-group" @click.stop="removeCombination(index)">
                      <el-button type="danger" link>删除组合</el-button>
                    </div>
                  </div>
                </div>

                <!-- 添加新组合按钮 -->
                <el-button v-if="product.slider_items.length < 5" class="add-combination" @click="addCombination">
                  + 添加内容组合（剩余 {{ 5 - product.slider_items.length }}）
                </el-button>
              </div>
            </el-form-item>

            <el-form-item label="商品价格">
              <el-input v-model.number="product.price" placeholder="请输入商品价格" />
            </el-form-item>
            <el-form-item label="商品库存">
              <el-input v-model.number="product.stock" placeholder="请输入商品库存" />
            </el-form-item>
            <el-form-item label="商品描述">
              <el-input v-model="product.description" placeholder="请输入商品描述" type="textarea" />
            </el-form-item>
            <el-form-item label="商品图片">
              <el-input v-model="product.picture" placeholder="请输入商品图片链接" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="addProduct">添加</el-button>
              <el-button @click="resetForm">重置</el-button>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
    </el-main>
  </el-container>

</template>

<script setup>
import Functions from '@/components/merchant/Functions.vue';
import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';
import { ElMessage } from 'element-plus';
import axios from 'axios';

const router = useRouter();
const store = useStore();
const addProductFormRef = ref(null);

// Vuex 数据
const categories = computed(() => store.getters['category/categories']);


const product = ref({
  name: '',
  categories: [],
  price: 0,
  stock: 0,
  description: '',
  img_url: '',
  // 修复初始数据结构不一致问题
  slider_items: []  // 替换原来的 slider_imgs: [[]]
});


const resetForm = () => {
  product.value = {
    name: '',
    categories: [],
    price: 0,
    stock: 0,
    description: '',
    img_url: '',
    slider_items: []  // 保持数据结构一致
  };
  activeSliderIndex.value = 0;
};


const isVideo = (url) => {
  return url?.match(/\.(mp4|mov|avi)$/i);
};

const addProduct = async () => {
  try {
    console.log('ProductDetail:', product)
    const payload = {
      ...product.value,
      categories: product.value.categories.map(id => ({
        id,
        name: categories.value.find(c => c.id === id)?.name
      }))
    };
    
    await axios.post('/api/merchant/products', payload);
    ElMessage.success('商品添加成功');
    router.push('/merchant/products');
  } catch (error) {
    ElMessage.error('商品添加失败：' + error.message);
  }
};


const beforeAvatarUpload = (file) => {
  const isImage = ['image/jpeg', 'image/png'].includes(file.type);
  const isLt2M = file.size / 1024 / 1024 < 2;

  if (!isImage) {
    ElMessage.error('只支持 JPG/PNG 格式!');
    return false;
  }
  if (!isLt2M) {
    ElMessage.error('图片大小不能超过 2MB!');
    return false;
  }
  return true;
};

// 在 script setup 中添加缺失的方法
const removeSlider = (index) => {
  product.value.slider_imgs.splice(index, 1)
}

const handleSliderUpload = (response) => {
  if (response.url) {
    if (product.value.slider_imgs.length >= 5) {
      ElMessage.warning('最多只能上传5个内容')
      return
    }
    product.value.slider_imgs.push(response.url)
  }
}

const beforeSliderUpload = () => {
  if (product.value.slider_imgs.length >= 5) {
    ElMessage.error('轮播图最多只能上传5个内容')
    return false
  }
  return true
}

// 修正后的 handleMainUpload 方法
const handleMainUpload = (response) => {
  if (response.url) {
    product.value.img_url = response.url
    // 自动插入到轮播图首位并去重
    product.value.slider_imgs = [
      response.url,
      ...product.value.slider_imgs.filter(url => url !== response.url)
    ].slice(0, 5)
    ElMessage.success('主图上传成功')
  } else {
    ElMessage.error('上传失败')
  }
}

const imageList = computed(() => product.value.slider_imgs.filter(item => !isVideo(item)));
const videoList = computed(() => product.value.slider_imgs.filter(item => isVideo(item)));


const beforeImageUpload = (file) => {
  const isImage = ['image/jpeg', 'image/png'].includes(file.type);
  if (!isImage) {
    ElMessage.error('只支持 JPG/PNG 格式!');
    return false;
  }
  return true;
};

const beforeVideoUpload = (file) => {
  const isVideo = file.type === 'video/mp4';
  const isLt50M = file.size / 1024 / 1024 < 50;
  
  if (!isVideo) {
    ElMessage.error('只支持 MP4 格式!');
    return false;
  }
  if (!isLt50M) {
    ElMessage.error('视频大小不能超过 50MB!');
    return false;
  }
  return true;
};

// 添加组合条目方法
const addCombination = () => {
  if (product.value.slider_items.length < 5) {
    product.value.slider_items.push({ image: null, video: null });
  }
};

// 修改后的上传处理方法
const handleImageUpload = (response, index) => {
  if (response.url) {
    product.value.slider_items[index].image = response.url;
  }
};

const handleVideoUpload = (response, index) => {
  if (response.url) {
    product.value.slider_items[index].video = response.url;
  }
};

// 删除媒体内容方法
const removeMedia = (index, type) => {
  product.value.slider_items[index][type] = null;
};

// 删除整个组合
const removeCombination = (index) => {
  product.value.slider_items.splice(index, 1);
};
</script>

// 新增样式调整
<style scoped>
.combination-item {
  margin-bottom: 16px;
  border: 1px solid #ebeef5;
  border-radius: 8px;
  padding: 12px;
}

.combination-content {
  display: flex;
  align-items: center;
  gap: 20px;
}

.media-group {
  flex: 1;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 15px;
}

.remove-group {
  flex-shrink: 0;
  width: 100px;
  text-align: center;
}

/* 保持其他原有样式不变... */
</style>