package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"linkShortener/internal/pkg/shortLink"
	"linkShortener/internal/pkg/shortLink/mocks"
	"testing"
)

var alphabet = []byte("_0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lenShortenLink = 10

func TestCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRep := mocks.NewMockShortLinkRepository(mockCtrl)
	testUc := ShortLinkUC{
		mockRep,
		alphabet,
		lenShortenLink,
	}

	originalLink := "originalLink"
	shortenLink := "shortenLink"

	mockRep.EXPECT().Exists(originalLink).Return(shortenLink, nil).Times(1)
	sl, err := testUc.Create(originalLink)
	require.NoError(t, err)
	assert.Equal(t, shortenLink, sl)

	gomock.InOrder(
		mockRep.EXPECT().Exists(originalLink).Return("", shortLink.LinkError{"link not found"}).Times(1),
		mockRep.EXPECT().Create(originalLink, gomock.Any()).Return(nil).Times(1),
		)
	_, err = testUc.Create(originalLink)
	require.NoError(t, err)

	gomock.InOrder(
		mockRep.EXPECT().Exists(originalLink).Return("", shortLink.LinkError{"link not found"}).Times(1),
		mockRep.EXPECT().Create(originalLink, gomock.Any()).Return(shortLink.LinkError{"short link exists"}).Times(1),
		mockRep.EXPECT().Create(originalLink, gomock.Any()).Return(nil).Times(1),
	)
	_, err = testUc.Create(originalLink)
	require.NoError(t, err)

	gomock.InOrder(
		mockRep.EXPECT().Exists(originalLink).Return("", shortLink.LinkError{"link not found"}).Times(1),
		mockRep.EXPECT().Create(originalLink, gomock.Any()).Return(errors.New("db error")).Times(1),
	)
	_, err = testUc.Create(originalLink)
	require.Error(t, err)
}

func TestGet(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRep := mocks.NewMockShortLinkRepository(mockCtrl)
	testUc := ShortLinkUC{
		mockRep,
		alphabet,
		lenShortenLink,
	}

	originalLink := "originalLink"
	shortenLink := "shortenLink"

	mockRep.EXPECT().Get(shortenLink).Return(originalLink, nil).Times(1)
	ol, err := testUc.Get(shortenLink)
	require.NoError(t, err)
	assert.Equal(t, originalLink, ol)
}