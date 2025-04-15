package util

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// DownloadFromCOS 从腾讯云COS下载文件到指定目录
// fileURL: 文件URL地址
// saveDir: 保存目录（为空则保存到当前目录）
// 返回: 下载文件的完整路径和错误信息
func DownloadFromCOS(fileURL, saveDir string) (string, error) {
	resp, err := http.Get(fileURL)
	if err != nil {
		return "", fmt.Errorf("下载请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}

	// 获取文件名
	fileName, err := getFileName(resp)
	if err != nil {
		return "", fmt.Errorf("获取文件名失败: %w", err)
	}

	// 处理保存目录
	if saveDir == "" {
		saveDir = "." // 当前目录
	}

	// 创建目录（如果不存在）
	if err := os.MkdirAll(saveDir, 0755); err != nil {
		return "", fmt.Errorf("创建目录失败: %w", err)
	}

	// 构建完整保存路径
	fullPath := filepath.Join(saveDir, fileName)

	// 创建本地文件
	file, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("文件创建失败: %w", err)
	}
	defer file.Close()

	// 流式写入文件
	if _, err := io.Copy(file, resp.Body); err != nil {
		return "", fmt.Errorf("文件写入失败: %w", err)
	}

	return fullPath, nil
}

// getFileName 获取文件名
func getFileName(resp *http.Response) (string, error) {
	disposition := resp.Header.Get("Content-Disposition")
	if disposition == "" {
		return "", fmt.Errorf("Content-Disposition头不存在")
	}

	_, params, err := mime.ParseMediaType(disposition)
	if err != nil {
		return "", fmt.Errorf("解析Content-Disposition失败: %w", err)
	}

	// 优先使用UTF-8编码的文件名
	if name := params["filename*"]; name != "" {
		if strings.Contains(name, "''") {
			parts := strings.SplitN(name, "''", 2)
			if decoded, err := url.QueryUnescape(parts[1]); err == nil {
				return decoded, nil
			}
		}
	}

	// 使用普通文件名
	if name := params["filename"]; name != "" {
		return strings.Trim(name, `"`), nil
	}

	return "", fmt.Errorf("无法从响应头中提取文件名")
}
